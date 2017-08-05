// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// apidata: Processes requests to Youtube API, downloads video streams

package api

import (	
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	
	"github.com/Sirupsen/logrus"

	"google.golang.org/api/youtube/v3"
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



func printVideosListResults(response *youtube.VideoListResponse) {
        for _, item := range response.Items {
                fmt.Println(item.Id, ": ", item.Snippet.Title)
        }
}

//Prints the video list by ID.
func videosListById(service *youtube.Service, part string, id string) {
        call := service.Videos.List(part)
        if id != "" {
                call = call.Id(id)
        }
        response, err := call.Do()
        handleError(err, "")
        printVideosListResults(response)
}



//Gets Video Data from Youtube URL
func APIGetVideoStream(service youtube.Service, url string)(videoData []byte, err error) {

	videoStream := new(RawVideoData)
	
	//Gets video Id
	id , err := getVideoId(url)
	auth.HandleError(err, "Invalid youtube URL.")
	
	//Get Video Data stream
	videoUrl := videoExtractor + id
	resp, er := http.Get(videoUrl)
	auth.HandleError(er, "Error in GET request)
	defer resp.Body.Close()
	out, e := ioutil.ReadAll(resp.Body)
	auth.HandleError(e, "Error reading video data")
	if err = json.Unmarshal(out, &a.output); err != nil {
		logrus.Errorf("Error JSON Unmarshall: %v", err)
	}
	//Extract Video information.
	videoInfo := videosListById(service, "snippet,contentDetails", id)//fileDetails part not permitted.
	
	//Get Data stream from video response
	if err = json.Unmarshal(out, &videoStream); err != nil {
		logrus.Errorf("Error JSON Unmarshall: %v", err)
	}
	
	//Download data stream to memory.
	
	//convert video file to flv or mp3


}



func ApiDownloadVideo() {


}



func decodeVideoInfo(response string) (streams streamList, err error) {
	// decode

	answer, err := url.ParseQuery(response)
	if err != nil {
		err = fmt.Errorf("parsing the server's answer: '%s'", err)
		return
	}

	// check the status

	err = ensureFields(answer, []string{"status", "url_encoded_fmt_stream_map", "title", "author"})
	if err != nil {
		err = fmt.Errorf("Missing fields in the server's answer: '%s'", err)
		return
	}

	status := answer["status"]
	if status[0] == "fail" {
		reason, ok := answer["reason"]
		if ok {
			err = fmt.Errorf("'fail' response status found in the server's answer, reason: '%s'", reason[0])
		} else {
			err = errors.New(fmt.Sprint("'fail' response status found in the server's answer, no reason given"))
		}
		return
	}
	if status[0] != "ok" {
		err = fmt.Errorf("non-success response status found in the server's answer (status: '%s')", status)
		return
	}

	log("Server answered with a success code")

	/*
	for k, v := range answer {
		log("%s: %#v", k, v)
	}
	*/

	// read the streams map

	stream_map := answer["url_encoded_fmt_stream_map"]

	// read each stream

	streams_list := strings.Split(stream_map[0], ",")

	log("Found %d streams in answer", len(streams_list))

	for stream_pos, stream_raw := range streams_list {
		stream_qry, err := url.ParseQuery(stream_raw)
		if err != nil {
			log(fmt.Sprintf("An error occured while decoding one of the video's stream's information: stream %d: %s\n", stream_pos, err))
			continue
		}
		err = ensureFields(stream_qry, []string{"quality", "type", "url"})
		if err != nil {
			log(fmt.Sprintf("Missing fields in one of the video's stream's information: stream %d: %s\n", stream_pos, err))
			continue
		}
		/* dumps the raw streams
		log(fmt.Sprintf("%v\n", stream_qry))
		*/
		stream := stream{
			"quality": stream_qry["quality"][0],
			"type": stream_qry["type"][0],
			"url": stream_qry["url"][0],
			"sig": "",
			"title": answer["title"][0],
			"author": answer["author"][0],
		}
		
		if sig, exists := stream_qry["sig"]; exists { // old one
			stream["sig"] = sig[0]
		}
		
		if sig, exists := stream_qry["s"]; exists { // now they use this
			stream["sig"] = sig[0]
		}
		
		streams = append(streams, stream)

		quality := stream.Quality()
		if quality == QUALITY_UNKNOWN {
			log("Found unknown quality '%s'", stream["quality"])
		}

		format := stream.Format()
		if format == FORMAT_UNKNOWN {
			log("Found unknown format '%s'", stream["type"])
		}

		log("Stream found: quality '%s', format '%s'", quality, format)
	}

	log("Successfully decoded %d streams", len(streams))

	return
}
