package main

// #cgo LDFLAGS: ./pkg/libydl.a -ldl
// #include "./pkg/download.h"
import "C"

import (
	"flag"
	"log"
	"os"
	"strings"
)

const (

	// help Banner
	BANNER = `` +
		`ydl - simple youtube downloader` + "\n\n" +
		`Usage: [OPTIONS] [ARGS]` + "\n" +
		"\t" + `ydl -id video url or id` + "\n" +
		"\t" + `ydl -path download path (defaults to '.')` + "\n" +
		"\t" + `Example: ydl -id https://www.youtube.com/watch?v=lWEbEtr_Vng` + "\n\n\n"

	// current version
	VERSION = "v1.0"

	// default maximum concurrent downloads
	MAXDOWNLOADS = 5
)

var (
	// Command line flags
	ids     string
	version bool
	format  string
	path    string
	bitrate uint
)

func init() {
	// parse flags
	flag.StringVar(&ids, "id", "", "video url or video id; separate multiple ids with a comma.")
	flag.StringVar(&path, "path", ".", "download file path")

	flag.Usage = func() {
		log.Fatalf("%s \t %s", BANNER, VERSION)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	args := os.Args
	if len(args) < 2 {
		usageAndExit(BANNER, 2)
	}
	if path == "" {
		path, _ = os.Getwd()
	}

	// parse urls
	urls := parseUrls(ids)

	// start download
	if err := concurrentDownload(MAXDOWNLOADS, format, urls); err != nil {
		log.Fatalf("Unable to download video(s): %v with errors => %v", urls, err)
	}
}

// parseUrls for video download
func parseUrls(urls string) []string {
	if ids == "" {
		return []string{os.Args[1]}
	} else {
		return strings.Split(ids, ",")
	}
}

//DownloadStreams download a batch of elements asynchronously
func concurrentDownload(maxOperations int, format string, urls []string) error {
	for _, url := range urls {
		// download video
		go func(url string) {
			cUrl := C.CString(url)
			cPath := C.CString(path)

			C.download(cUrl, cPath)
		}(url)
	}
	return nil
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		log.Fatalf(message)
	}
	flag.Usage()
	os.Exit(exitCode)
}
