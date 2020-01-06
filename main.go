// main program entry
package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const (

	// help Banner
	BANNER = `` +
		`youtube-dl - Simple youtube video/audio DownloadStreams` + "\n\n" +
		`Usage: youtube-dl [OPTIONS] [ARGS]` + "\n"

	// current version
	VERSION = "v0.2"

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

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func init() {
	// parse flags
	flag.StringVar(&ids, "id", "", "video url or video id; separate multiple ids with a comma.")
	flag.StringVar(&format, "format", "", "download file format(mp3 or flv)")
	flag.StringVar(&path, "path", ".", "download file path")
	flag.BoolVar(&version, "version", false, "print version number")
	flag.UintVar(&bitrate, "bitrate", 192, "audio bitrate")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf("%s \t %s", BANNER, VERSION))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			logrus.Fatalf("%v", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	runtime.SetBlockProfileRate(20)

	args := os.Args
	if len(args) < 2 {
		usageAndExit(BANNER, 2)
	}
	if path == "" {
		path, _ = os.Getwd()
	}

	urls := parseUrls(ids)
	beginDownload(urls)
}

// parseUrls for video download
func parseUrls(urls string) []string {
	if ids == "" {
		return []string{os.Args[1]}
	} else {
		return strings.Split(ids, ",")
	}
}

func beginDownload(urls []string) {

	if len(urls) < 2 {
		if vId, err := getVideoId(urls[0]); err != nil {
			logrus.Errorf("Error fetching videoId: %v", err)
		} else {
			if err := decodeVideoStream(vId, path, format, bitrate); err != nil {
				logrus.Errorf("Unable to beginDownload: %v", err)
			}
		}
	} else {
		if err := concurrentDownload(MAXDOWNLOADS, format, path, bitrate, urls); err != nil {
			logrus.Errorf("Unable to concurrently download videos: %v with errors => %v", urls, err)
		}
	}
}

//DownloadStreams download a batch of elements asynchronously
func concurrentDownload(maxOperations int, format, outputPath string, bitrate uint, urls []string) <-chan error {

	var wg sync.WaitGroup
	wg.Add(len(urls))

	ch := make(chan error, maxOperations)
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()

			if videoId, err := getVideoId(url); err != nil {
				ch <- err
			} else {
				ch <- decodeVideoStream(videoId, path, format, bitrate)
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
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
