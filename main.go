// Copyright 2020 YTD Authors. All rights reserved.
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
	"strings"

	"github.com/Ch3ck/youtube-dl/api"
	"github.com/Sirupsen/logrus"
)

const (

	//BANNER for ytd which prints the help info
	BANNER = `
		youtube-dl -h
		youtube-dl - Simple youtube video/audio downloader

		Usage: youtube-dl [OPTIONS] [ARGS]
	`
	//VERSION which prints the ytd version.
	VERSION = "v0.2"
)

var (
	ids     string
	version bool
	format  string
	path    string
	bitrate uint
)

const (
	defaultMaxDownloads = 5
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func init() {
	// parse flags
	flag.StringVar(&ids, "id", "", "Youtube Video IDs. Separated then by using a comma.")
	flag.StringVar(&format, "format", "", "File Format(mp3, webm, flv)")
	flag.StringVar(&path, "path", ".", "Output Path")
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.UintVar(&bitrate, "bitrate", 192, "Audio Bitrate")

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
	if ids == "" {
		url := os.Args[1]
		startProcessing([]string{url})
	} else {
		startProcessing(strings.Split(ids, ","))
	}
}

func startProcessing(urls []string) {
	ch := api.DownloadStreams(defaultMaxDownloads, format, path, bitrate, urls)
	for err := range ch {
		//Extract Video data and decode
		if err != nil {
			logrus.Errorf("Error decoding Video stream: %v", err)
		}
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
