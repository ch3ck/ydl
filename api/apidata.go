// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// apidata: Processes requests to Youtube API, downloads video streams

package api

import (	
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

	video := new(RawVideoData)//raw video data
	var decodedVideo []string //decoded video data
	
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
	
	output, er := url.ParseQuery(out)
	if e != nil {
		logrus.Fatalf("Error Parsing video byte stream", e)
	}
	//fmt.Println(string(output))
	
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
	
	
	//Download data stream to memory and convert to proper format
	//NOTE: Use ffmpeg go bindings for this use case.
	
}



func APIDownloadVideo(videoStream map[string][]string) ([]byte, err) {
	func (stream stream) download(out io.Writer) error {
	url := stream.Url()

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


//Search youtube API for video data.
func SearchApi(service youtube.Service) {

	
	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	printIDs("Videos", videos)
	printIDs("Channels", channels)
	printIDs("Playlists", playlists)
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func APIPrintIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
