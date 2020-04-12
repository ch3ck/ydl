package main

import (
	"context"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
	"github.com/wader/goutubedl"
)

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

// decodeVideoStream processes downloaded video stream and
// decodeVideoStream calls helper functions and writes the
// output in the required format
func decodeVideoStream(videoUrl, format string) error {

	// Get video data
	res, err := goutubedl.New(context.Background(), videoUrl, goutubedl.Options{})
	if err != nil {
		logrus.Errorf("Unable to create goutube object %s: %v", videoUrl, err)
		return err
	}

	file := removeWhiteSpace(res.Info.Title) + fixExtension(format)
	videoStream, err := res.Download(context.Background(), "best")
	if err != nil {
		logrus.Errorf("Unable to download %s stream: %v", format, err)
		return err
	}
	defer videoStream.Close()

	// Create output file
	fp, err := os.OpenFile(file, os.O_CREATE, 0755)
	if err != nil {
		logrus.Errorf("Unable to create output file: %v", err)
		return err
	}
	defer fp.Close()

	io.Copy(fp, videoStream)

	return nil
}