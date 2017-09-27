// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// ytd program entry.

package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/Ch3ck/youtube-dl/api"
	"github.com/Sirupsen/logrus"
)

const (

	//BANNER for ytd which prints the help info
	BANNER = "ytd -id 'videoId' -format mp3 -bitrate 123  -path ~/Downloads/ videoUrl %s\n"
	//VERSION which prints the ytd version.
	VERSION = "v0.1"
)

var (
	id      string
	version bool
	format  string
	path    string
	bitrate uint
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

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
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			logrus.Fatalf("%v", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	runtime.SetBlockProfileRate(20)

	if path == "" {
		path, _ = os.Getwd()
	}
	if len(os.Args) == 1 {
		usageAndExit(BANNER, -1)
	}
	//Get Video Id
	if id == "" {
		url := os.Args[1]
		ID, _ = api.GetVideoId(url)
	} else {
		ID, _ = api.GetVideoId(id)
	}

	//Extract Video data and decode
	if err := api.APIGetVideoStream(format, ID, path, bitrate); err != nil {
		logrus.Errorf("Error decoding Video stream: %v", err)
	}
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}
