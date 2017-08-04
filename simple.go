//curl https://youtube.com/get_video_info?video_id=xbqyAd_sZG0
//url := "http://youtube.com/get_video_info?video_id="
// https://www.googleapis.com/youtube/v3/captions/id

package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

const developerKey = "AIzaSyCZSy5sOGsZrOrI0vLtowf_VJ-tl_USzNE"

func main() {

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make GET request to Youtube API.
	call := service.Videos.List("snippet, recordingDetails")
	call.Id("Ks-_Mh1QhMc")
	resp, err := call.Do()
	if err != nil {
		log.Fatalf("Error getting Video response: %v", err)
	}
	for cnt, item := range resp.Items {
                fmt.Printf("\n %d: %+v\n", cnt, item)
    }
    
	return
}
