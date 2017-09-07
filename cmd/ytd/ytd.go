// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// ytd program entry.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/Ch3ck/ytd/api"
	"github.com/Ch3ck/ytd/auth"
	"github.com/Sirupsen/logrus"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const (

	//Developer key
	devKey = "" //Generated from OAuth

	//BANNER for ytd which prints the help info
	BANNER = "ytd - %s\n"
	//VERSION which prints the ytd version.
	VERSION = "v0.1"
)

var (
	key     string
	query   string
	version bool

	//TODO: Currently only the first result will be returned on CLI
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
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

func init() {
	// parse flags
	flag.StringVar(&key, "key", "", "Youtube API key")
	flag.StringVar(&query, "query", "", "Youtube search Query")

	flag.BoolVar(&version, "version", false, "print version and exit")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if version {
		logrus.Infof("%s", VERSION)
		os.Exit(0)
	}

	// set log level
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

}

func main() {
	ctx := context.Background()
	service, err = auth.CreateYoutubeService(ctx)
	auth.HandleError(err, "Error creating YouTube client")

	//channelsListByUsername(service, "snippet,contentDetails,statistics", "GoogleDevelopers")
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
