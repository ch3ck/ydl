// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// ytd program entry.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ch3ck/ytd/api"
	"github.com/Sirupsen/logrus"
)

const (

	//BANNER for ytd which prints the help info
	BANNER = "ytd -id 'videoId' -format mp3 -bitrate 123  -path ~/Downloads/ videoUrl%s\n"
	//VERSION which prints the ytd version.
	VERSION = "v0.1"
)

var (
	id      string
	version bool
	format  string
	path    string
	bitrate uint
	file    string
)

func init() {
	// parse flags
	flag.StringVar(&id, "id", "", "Youtube Video ID")
	flag.StringVar(&format, "format", "", "File Format(mp3, webm, flv)")
	flag.StringVar(&path, "path", ".", "Output Path")
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.UintVar(&bitrate, "bitrate", 123, "Audio Bitrate")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if version {
		logrus.Infof("%s", VERSION)
		os.Exit(0)
	}
}

func main() {
	var ID string
	var rawVideo api.RawVideoData
	if len(os.Args) == 1 {
		usageAndExit(BANNER, -1)
	}

	//Get Video Id
	if id == "" {
		url := os.Args[1]
		ID, _ = api.GetVideoId(id)
	} else {
		ID, _ = api.GetVideoId(id)
	}

	//Extract Video data and decode
	video, err := api.APIGetVideoStream(ID, rawVideo)
	if err != nil {
		logrus.Errorf("Error decoding Video stream: %v", err)
	}

	//Convert and Download video data
	//create output file name and set path properly.
	file = path + video["title"] + video["author"]
	if format == "mp3" {
		file = file + ".mp3"

	} else { //defaults to flv format for video files.)
		file = file + ".flv"
	}

	err = api.APIConvertVideo(file, bitrate, ID, video)
	if err != nil {
		logrus.Errorf("Error downloading video: %v", err)
	}
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		logrus.Infof(os.Stderr, message)
		logrus.Infof(os.Stderr, "\n\n")
	}
	flag.Usage()
	logrus.Infof(os.Stderr, "\n")
	os.Exit(exitCode)
}
