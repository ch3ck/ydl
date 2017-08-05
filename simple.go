//curl https://youtube.com/get_video_info?video_id=xbqyAd_sZG0
//url := "http://youtube.com/get_video_info?video_id="
// https://www.googleapis.com/youtube/v3/captions/id

package main

import (
	//"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"io/ioutil"
	//"log"
	"net/http"
	"net/url"
	
	"github.com/Sirupsen/logrus"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)


const developerKey = "AIzaSyCZSy5sOGsZrOrI0vLtowf_VJ-tl_USzNE"

type RawVideoData struct {
	Title     				string `json:"title"`
	Author     				string `json:"author`
	Status                 	string `json:"status"`
	URLEncodedFmtStreamMap 	map[string][]string `json:"url_encoded_fmt_stream_map"`
}


func main() {

	videoExtractor := "https://youtube.com/get_video_info?video_id="
	var decodedVideo []string
	id := "Ks-_Mh1QhMc"
	
	video := new(RawVideoData)
	
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		logrus.Fatalf("Error creating new YouTube client: %v", err)
	}

    
	logrus.Infof("Starting the Video extration process.")
	
	//Get Video Data stream
	videoUrl := videoExtractor + id
	res, er := http.Get(videoUrl)
	if res.StatusCode != 200 || er != nil {
		logrus.Infof("Status response: %v. Err: %v", res.StatusCode, er)
	}
	defer res.Body.Close()
	out, e := ioutil.ReadAll(res.Body)
	if e != nil {
		logrus.Fatalf("Error reading video data", e)
	}
	output, er := url.ParseQuery(out)
	if e != nil {
		logrus.Fatalf("Error unmarshalling output", e)
	}
	fmt.Println(string(output))
	
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
	
	//Download video
	

    
	return
}


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





