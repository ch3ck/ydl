// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// apidata: Processes requests to Youtube API, downloads video streams

package api

import (	
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	
	"github.com/Sirupsen/logrus"
)


//const variables
const (
	
	//Video extractor
	videoExtractor = "http://youtube.com/get_video_info?video_id="
)

//Youtube Downloader Data file.
type RawVideoData struct {
	Title     				string `json:"title"`
	Author     				string `json:"author`
	Status                 	string `json:"status"`
	URLEncodedFmtStreamMap 	map[string][]string `json:"url_encoded_fmt_stream_map"`
}
		
}
type ApiData struct {
	FileName    string
	Title       string
	description string
	category    string
	keywords    string
	privacy     string
	DataStream  []byte
}


//gets the Video ID from youtube url
func getVideoId(url string) ( string, error) {

	s := strings.Split(url, "?v=")
	s = strings.Split(s[1], "&")
	if len(s[0]) == 0 {
		return s[0], errors.New("Empty string")
	}
	
	return s[0], nil
}


//Gets Video Info, Decode Video Info from a Video ID.
func APIGetVideoStream(url string)(videoData []byte, err error) {

	video := new(RawVideoData)//raw video data
	var decodedVideo []string //decoded video data
	
	//Gets video Id
	id , err := getVideoId(url)
	
	//Get Video Data stream
	videoUrl := videoExtractor + id
	resp, er := http.Get(videoUrl)
	if er != nil {
		logrus.Errorf("Error in GET request: %v", er)
	}
	
	defer resp.Body.Close()
	out, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		logrus.Errorf("Error reading video data: %v", e)
	}
		
	output, er := url.ParseQuery(out)
	if e != nil {
		logrus.Fatalf("Error Parsing video byte stream: %v", e)
	}
	
	//Process Video stream
	video.URLEncodedFmtStreamMap = output["url_encoded_fmt_stream_map"]
	video.Author  = output["author"]
	video.Title = output["title"]
	video.Status = output["status"]
	
	//Decode Video
	outputStreams := strings.Split(video.URLEncodedFmtStreamMap[0], ",")
	for cur, raw_data := range outputStream {
		//decoding raw data stream
		dec_data, err := url.ParseQuery(raw_data)
		if err != nil {
			logrus.Errorf("Error Decoding Video data: %d, %v", cur, err)
			continue
		}
		
		data := map[string]string{
			"quality": dec_data["quality"][0],
			"type": dec_data["type"][0],
			"url": dec_data["url"][0],
			"sig": dec_data["sig"][0],
			"title": video.Title,
			"author": video.Author,
			"format": dec_data["format"][0],
		}
		
		decodedVideo = append(decodedVideo, data)
		logrus.Infof("\nDecoded %d bytes of '%s", in '%s' format, len(decodedVideo), dec_data["quality"][0], dec_data["format"][0])
	}
	
	return decodedVideo
}


// 
//Downloads decoded video stream.
func APIDownloadVideo(videoUrl string) error {

	log("Downloading stream from '%s'", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("requesting stream: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("reading answer: non 200 status code received: '%s'", err)
	}
	length, err := io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("saving file: %s (%d bytes copied)", err, length)
	}

	log("Downloaded %d bytes", length)

	return nil
}
