package main

import (
	"context"
	"io"
	"os"
	"os/user"
	"path/filepath"
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
func decodeVideoStream(videoUrl, path, format string) error {

	// Get video data
	res, err := goutubedl.New(context.Background(), videoUrl, goutubedl.Options{})
	if err != nil {
		logrus.Errorf("Unable to create goutube object %s: %v", videoUrl, err)
	}

	file := removeWhiteSpace(res.Info.Title) + fixExtension(format)
	videoStream, err := res.Download(context.TODO(), format)
	if err != nil {
		logrus.Errorf("Unable to download %s stream: %v", format, err)
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

	io.Copy(fp, videoStream)
	videoStream.Close()

	return nil
}
