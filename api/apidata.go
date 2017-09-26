// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// apidata: Processes requests to Youtube API, downloads video streams

package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	"github.com/Sirupsen/logrus"
)

//const variables
const (
	audioBitRate = 123 //default audio bit rate.

	//Video extractor
	videoExtractor = "https://youtube.com/get_video_info?video_id="
)

type stream map[string]string

//Youtube Downloader Data file.
type RawVideoData struct {
	Title                  string   `json:"title"`
	Author                 string   `json:"author"`
	Status                 string   `json:"status"`
	URLEncodedFmtStreamMap []stream `json:"url_encoded_fmt_stream_map"`
	VideoId                string
	VideoInfo              string
	Vlength                float64
	dpercent               chan int64
}

func (v *RawVideoData) Write(b []byte) (n int, err error) {
	n = len(b)
	totalbytes, dlevel := 0.0, 0.0
	v.Vlength = totalbytes + float64(n)
	curPercent := ((totalbytes / v.Vlength) * 100)
	if (dlevel <= curPercent) && (dlevel < 100) {
		dlevel++
		v.dpercent <- int64(dlevel)
	}
	return
}

//gets the Video ID from youtube url
func GetVideoId(url string) (string, error) {
	if !strings.Contains(url, "youtube.com") {
		return "", errors.New("Invalid Youtube link")
	}
	s := strings.Split(url, "?v=")
	s = strings.Split(s[1], "&")
	if len(s[0]) == 0 {
		return s[0], errors.New("Empty string")
	}

	return s[0], nil
}

//Gets Video Info, Decode Video Info from a Video ID.
func APIGetVideoStream(format, id, path string, bitrate uint) (err error) {

	video := new(RawVideoData) //raw video data
	var decodedVideo []string  //decoded video data

	//Get Video Data stream
	video.VideoId = id
	videoUrl := videoExtractor + id
	video.VideoInfo = videoUrl
	resp, er := http.Get(videoUrl)
	if er != nil {
		logrus.Errorf("Error in GET request: %v", er)
	}

	defer resp.Body.Close()
	out, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		logrus.Errorf("Error reading video data: %v", e)
	}

	output, er := url.ParseQuery(string(out))
	if e != nil {
		logrus.Errorf("Error parsing video byte stream: %v", e)
		return nil
	}

	status, ok := output["status"]
	if !ok {
		err = fmt.Errorf("No response status in server")
		return err
	}
	if status[0] == "fail" {
		reason, ok := output["reason"]
		if ok {
			err = fmt.Errorf("'fail' response status found in the server, reason: '%s'", reason[0])
		} else {
			err = errors.New(fmt.Sprint("'fail' response status found in the server, no reason given"))
		}
		return err
	}
	if status[0] != "ok" {
		err = fmt.Errorf("non-success response status found in the server (status: '%s')", status)
		return err
	}

	// read the streams map
	video.Author = output["author"][0]
	video.Title = output["title"][0]
	StreamMap, ok := output["url_encoded_fmt_stream_map"]
	if !ok {
		err = fmt.Errorf("Error reading encoded stream map.")
		return err
	}

	// read and decode streams.
	streamsList := strings.Split(string(StreamMap[0]), ",")
	var streams []stream
	for streamPos, streamRaw := range streamsList {
		streamQry, err := url.ParseQuery(streamRaw)
		if err != nil {
			logrus.Infof("An error occured while decoding one of the video's streams: stream %d: %s\n", streamPos, err)
			continue
		}
		var sig string
		if _, exist := streamQry["sig"]; exist {
			sig = streamQry["sig"][0]
		}

		streams = append(streams, stream{
			"quality": streamQry["quality"][0],
			"type":    streamQry["type"][0],
			"url":     streamQry["url"][0],
			"sig":     sig,
			"title":   output["title"][0],
			"author":  output["author"][0],
		})
		logrus.Infof("Stream found: quality '%s', format '%s'", streamQry["quality"][0], streamQry["type"][0])
	}

	video.URLEncodedFmtStreamMap = streams
	//Download Video stream to file
	if format == "" {
		format = ".flv"
	} else {
		format = ".mp3"
	}
	//create output file name and set path properly.
	file := video.Title + format
	file = SpaceMap(file)
	vstream := streams[0]
	url := vstream["url"] + "&signature" + vstream["sig"]
	logrus.Infof("Downloading file to %s", file)
	if format == ".mp3" {
		err = ApiConvertVideo(file, id, format, bitrate, decodedVideo)
		if err != nil {
			logrus.Errorf("Error downloading audio: %v", err)
		}

	} else { //defaults to flv format for video files.)
		if err := ApiDownloadVideo(path, file, url, video); err != nil {
			logrus.Errorf("Error downloading video: %v", err)
		}
	}

	return nil
}

//remove whitespaces in filename
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
