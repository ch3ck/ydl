package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
	"github.com/viert/go-lame"
)

const (
	audioBitRate = 123

	streamApiUrl = "http://youtube.com/get_video_info?video_id="
)

type stream map[string]string

type RawVideoStream struct {
	VideoId                string
	VideoInfo              string
	Title                  string   `json:"title"`
	Author                 string   `json:"author"`
	URLEncodedFmtStreamMap []stream `json:"url_encoded_fmt_stream_map"`
	Status                 string   `json:"status"`
}

// removeWhiteSpace removes white spaces from string
// removeWhiteSpace returns a filename without whitespaces
func removeWhiteSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// fixExtension is a helper function that
// fixes file the extension
func fixExtension(str string) string {
	if strings.Contains(str, "mp3") {
		str = ".mp3"
	} else {
		str = ".flv"
	}

	return str
}

// encodeAudioStream consumes a raw data stream and
// encodeAudioStream encodes the data stream in mp3
func encodeAudioStream(file, path, surl string, bitrate uint) error {
	data, err := downloadVideoStream(surl)
	if err != nil {
		log.Printf("Http.Get\nerror: %s\nURL: %s\n", err, surl)
		return err
	}

	tmp, _ := os.OpenFile("_temp_", os.O_CREATE, 0755)
	defer tmp.Close()
	if _, err := tmp.Write(data); err != nil {
		logrus.Errorf("Failed to read response body: %v", err)
		return err
	}

	// Create output file
	currentDirectory, err := user.Current()
	if err != nil {
		logrus.Errorf("Error getting current user directory: %v", err)
		return err
	}

	outputDirectory := currentDirectory.HomeDir + "/Downloads/" + path
	outputFile := filepath.Join(outputDirectory, file)
	if err := os.MkdirAll(filepath.Dir(outputFile), 0775); err != nil {
		logrus.Errorf("Unable to create output directory: %v", err)
	}

	fp, err := os.OpenFile(outputFile, os.O_CREATE, 0755)
	if err != nil {
		logrus.Errorf("Unable to create output file: %v", err)
		return err
	}
	defer fp.Close()

	// write audio/video file to output
	reader := bufio.NewReader(tmp)
	writer := lame.NewEncoder(fp)
	defer writer.Close()

	writer.SetBrate(int(bitrate))
	writer.SetQuality(1)
	reader.WriteTo(writer)

	return nil
}

// encodeVideoStream consumes video data stream and
// encodeVideoStream encodes the video in flv
func encodeVideoStream(file, path, surl string) error {
	data, err := downloadVideoStream(surl)
	if err != nil {
		log.Printf("Http.Get\nerror: %s\nURL: %s\n", err, surl)
		return err
	}

	// Create output file
	currentDirectory, err := user.Current()
	if err != nil {
		logrus.Errorf("Error getting current user directory: %v", err)
		return err
	}

	outputDirectory := currentDirectory.HomeDir + "/Downloads/" + path
	outputFile := filepath.Join(outputDirectory, file)
	if err := os.MkdirAll(filepath.Dir(outputFile), 0775); err != nil {
		logrus.Errorf("Unable to create output directory: %v", err)
	}

	fp, err := os.OpenFile(outputFile, os.O_CREATE, 0755)
	if err != nil {
		logrus.Errorf("Unable to create output file: %v", err)
		return err
	}
	defer fp.Close()

	//saving downloaded file.
	if _, err = fp.Write(data); err != nil {
		logrus.Errorf("Unable to encode video stream: %s `->` %v", surl, err)
		return err
	}
	return nil
}

// downloadVideoStream downloads video streams from youtube
// downloadVideoStream returns the *http.Reponse body
func downloadVideoStream(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("Unable to fetch Data stream from URL(%s)\n: %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logrus.Errorf("Video Download error with status: '%v'", resp.StatusCode)
		return nil, errors.New("Non 200 status code received")
	}

	output, _ := ioutil.ReadAll(resp.Body)

	return output, nil
}

// getVideoId extracts the video id string from youtube url
// getVideoId returns a video id string to calling function
func getVideoId(url string) (string, error) {
	if len(url) < 15 {
		return url, nil
	} else {
		if !strings.Contains(url, "youtube.com") {
			return "", errors.New("Invalid Youtube URL")
		}

		s := strings.Split(url, "?v=")[1]
		if len(s) == 0 {
			return s, errors.New("Empty string")
		}

		return s, nil
	}
}

// decodeStream accept Values and decodes them individually
// decodeStream returns the final RawVideoStream object
func decodeStream(values url.Values, streams *RawVideoStream, rawstream []stream) error {
	streams.Author = values.Get("author")
	streams.Title = values.Get("title")
	streamMap := values.Get("url_encoded_fmt_stream_map")

	// read and decode streams
	streamsList := strings.Split(string(streamMap), ",")
	for streamPos, streamRaw := range streamsList {
		streamQry, err := url.ParseQuery(streamRaw)
		if err != nil {
			logrus.Infof("Error occured during stream decoding %d: %s\n", streamPos, err)
			continue
		}
		var sig string
		sig = streamQry.Get("sig")
		rawstream = append(rawstream, stream{
			"quality": streamQry.Get("quality"),
			"type":    streamQry.Get("type"),
			"url":     streamQry.Get("url"),
			"sig":     sig,
			"title":   values.Get("title"),
			"author":  values.Get("author"),
		})
		logrus.Infof("Stream found: quality '%s', format '%s'", streamQry.Get("quality"), streamQry.Get("type"))
	}

	streams.URLEncodedFmtStreamMap = rawstream
	return nil
}

// decodeVideoStream processes downloaded video stream and
// decodeVideoStream calls helper functions and writes the
// output in the required format
func decodeVideoStream(videoId, path, format string, bitrate uint) error {
	var decStreams []stream         //decoded video streams
	rawVideo := new(RawVideoStream) // raw video stream

	// Get video data
	rawVideo.VideoId = videoId
	rawVideo.VideoInfo = streamApiUrl + videoId

	data, err := downloadVideoStream(rawVideo.VideoInfo)
	if err != nil {
		logrus.Errorf("Unable to get video stream: %v", err)
		return err
	}

	parsedResp, err := url.ParseQuery(string(data))
	if err != nil {
		logrus.Errorf("Error parsing video byte stream: %v", err)
		return err
	}

	status, ok := parsedResp["status"]
	if !ok {
		return errors.New("No response from server")
	}

	reason, _ := parsedResp["reason"]
	if status[0] == "fail" {
		return errors.New(fmt.Sprintf("'fail' response with reason: %s", reason))
	} else if status[0] != "ok" {
		return errors.New(fmt.Sprintf("'non-success' response with reason: %s", reason))
	}

	if err := decodeStream(parsedResp, rawVideo, decStreams); err != nil {
		return errors.New("Unable to decode raw video streams")
	}

	file := removeWhiteSpace(rawVideo.Title) + fixExtension(format)
	surl := decStreams[0]["url"] + "&signature" + decStreams[0]["sig"]

	logrus.Infof("Downloading data to file: %s", file)
	if strings.Contains(file, "mp3") {
		if err := encodeAudioStream(file, path, surl, bitrate); err != nil {
			logrus.Errorf("Unable to encode %s: %v", format, err)
		}
	} else {
		if err := encodeVideoStream(file, path, surl); err != nil {
			logrus.Errorf("Unable to encode %s: %v", format, err)
		}
	}

	return nil
}
