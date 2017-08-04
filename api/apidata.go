// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// apidata: Processes requests to Youtube API, downloads video streams

package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	
	"github.com/Sirupsen/logrus"

	"google.golang.org/api/youtube/v3"
)

//Youtube Downloader Data file.
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

	s := strings.Split(url, "?v="))
	s = strings.Split(s[1], "&")
	if len(s[0]) == 0 {
		s[0], error.New("Empty string)
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

videosListById(service, "snippet,contentDetails,statistics", "Ks-_Mh1QhMc")




func APIGetVideoStream(service youtube.Service, url string)(videoData []byte, err error) {
	
	//Gets video Id
	id , err := getVideoId(url)
	auth.HandleError(err, "Invalid youtube URL.")
	
	//Get Video response stream
	videosListById(service, "snippet,contentDetails,liveStreamingDetails, fileDetails", id)//fileDetails part not permitted.
	
	//Get Data stream from video response
	
	
	//Download data stream to memory.
	
	//convert video file to flv or mp3




//retrieve uploads
func main() {
	flag.Parse()

	client, err := buildOAuthHTTPClient(youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Error building OAuth client: %v", err)
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	// Start making YouTube API calls.
	// Call the channels.list method. Set the mine parameter to true to
	// retrieve the playlist ID for uploads to the authenticated user's
	// channel.
	call := service.Channels.List("contentDetails").Mine(true)

	response, err := call.Do()
	if err != nil {
		// The channels.list method call returned an error.
		log.Fatalf("Error making API call to list channels: %v", err.Error())
	}

	for _, channel := range response.Items {
		playlistId := channel.ContentDetails.RelatedPlaylists.Uploads
		// Print the playlist ID for the list of uploaded videos.
		fmt.Printf("Videos in list %s\r\n", playlistId)

		nextPageToken := ""
		for {
			// Call the playlistItems.list method to retrieve the
			// list of uploaded videos. Each request retrieves 50
			// videos until all videos have been retrieved.
			playlistCall := service.PlaylistItems.List("snippet").
				PlaylistId(playlistId).
				MaxResults(50).
				PageToken(nextPageToken)

			playlistResponse, err := playlistCall.Do()

			if err != nil {
				// The playlistItems.list method call returned an error.
				log.Fatalf("Error fetching playlist items: %v", err.Error())
			}

			for _, playlistItem := range playlistResponse.Items {
				title := playlistItem.Snippet.Title
				videoId := playlistItem.Snippet.ResourceId.VideoId
				fmt.Printf("%v, (%v)\r\n", title, videoId)
			}

			// Set the token to retrieve the next page of results
			// or exit the loop if all results have been retrieved.
			nextPageToken = playlistResponse.NextPageToken
			if nextPageToken == "" {
				break
			}
			fmt.Println()
		}
	}
}

func ApiDownloadVideo() {


}main() {
	flag.Parse()

	if *filename == "" {
		log.Fatalf("You must provide a filename of a video file to upload")
	}

	client, err := buildOAuthHTTPClient(youtube.YoutubeUploadScope)
	if err != nil {
		log.Fatalf("Error building OAuth client: %v", err)
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       *title,
			Description: *description,
			CategoryId:  *category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: *privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(*keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(*keywords, ",")
	}

	call := service.Videos.Insert("snippet,status", upload)

	file, err := os.Open(*filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", *filename, err)
	}

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
}

