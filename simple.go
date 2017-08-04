//curl https://youtube.com/get_video_info?video_id=xbqyAd_sZG0
//url := "http://youtube.com/get_video_info?video_id="
// https://www.googleapis.com/youtube/v3/captions/id

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

const developerKey = "vqNWDovin1e8vR24HNdxF0_G"

func main() {

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make GET request to Youtube API.
	call := service.Videos.List("snippet, contentDetails, fileDetails,liveStreamingDetails")
	resp, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}
	for _, item := range response.Items {
                fmt.Println(item.Id, ": ", item.Snippet.Title)
        }
	body, err := ioutil.ReadAll(resp)
	if err != nil {
		fmt.Printf("Error reading output")
		return
	}

	fmt.Printf("\nGot %d bytes answer", len(body))
	return
}
