//curl https://youtube.com/get_video_info?video_id=xbqyAd_sZG0
//url := "http://youtube.com/get_video_info?video_id="
// https://www.googleapis.com/youtube/v3/captions/id

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	
	"github.com/Sirupsen/logrus"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

const developerKey = "AIzaSyCZSy5sOGsZrOrI0vLtowf_VJ-tl_USzNE"

func main() {

	videoExtractor := "https://youtube.com/get_video_info?video_id="
	var dat map[string]interface{}
	id := "Ks-_Mh1QhMc"
	
	
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		logrus.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make GET request to Youtube API.
	call := service.Videos.List("snippet, contentDetails, liveStreamingDetails, player")
	call.Id(id)
	resp, err := call.Do()
	if err != nil {
		logrus.Fatalf("Error getting Video response: %v", err)
	}
	for cnt, item := range resp.Items {
                fmt.Printf("\n %d: %+v\n", cnt, item)
    }
    
	logrus.Infof("Starting the Video extration process.")
	
	//Get Video Data stream
	videoUrl := videoExtractor + id
	res, er := http.Get(videoUrl)
	logrus.Fatalf("Error in GET request: %v", er)
	defer res.Body.Close()
	out, e := ioutil.ReadAll(res.Body)
	logrus.Fatalf("Error reading video data", e)
	if err = json.Unmarshal(out, &dat); err != nil {
		logrus.Errorf("Error JSON Unmarshall: %v", err)
	}
	
	fmt.Println(dat)
    
	return
}
