package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/kkdai/youtube/v2"
	"github.com/urfave/cli/v2"
)

const (
	VERSION = "v0.1.0"                                      // current version
	URL     = "https://www.youtube.com/watch?v=lWEbEtr_Vng" // default video url
	PATH    = "."                                           // default download path
)

func main() {
	app := &cli.App{
		Name:                 "ydl",
		Usage:                "simple youtube downloader",
		EnableBashCompletion: true,
		Version:              VERSION,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Nyah Check",
				Email: "hello@nyah.dev",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Value:    URL,
				Usage:    "Youtube video url or link",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "path",
				Value:   PATH,
				Aliases: []string{"p"},
				Usage:   "Destination for downloaded `FILE`",
			},
		},
		Action: func(c *cli.Context) error {
			// process args here
			cli.DefaultAppComplete(c)
			cli.HandleExitCoder(errors.New("invalid `ydl` command"))
			cli.ShowAppHelp(c)
			cli.ShowCompletions(c)
			cli.ShowVersion(c)

			// get app names
			fmt.Printf("Args: %#v\n", c.Args())
			fmt.Printf("IDs: %#v\n", c.String("id"))
			fmt.Printf("Path: %#v\n", c.String("path"))

			// parse urls
			url := c.String("id")
			path := c.String("path")

			// download files with go library
			return downloadVideo(url, path)

			// use rust library instead
			// return rsDownloadVideo(url, path)
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to download video(s) with error => %v", err)
	}
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

// downloadVideo downloads a video from a youtube url
// returns error if error occurs
func downloadVideo(url string, path string) error {
	client := youtube.Client{}

	videoInfo, err := client.GetVideo(url)
	if err != nil {
		return err
	}

	audioFormat := videoInfo.Formats.WithAudioChannels()
	videoStream, _, err := client.GetStream(videoInfo, &audioFormat[0])
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s.mp4", removeWhiteSpace(videoInfo.Title))
	fmt.Printf("\nFileName: %s", fileName)
	filePath := filepath.Join(path, fileName)
	videoFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer videoFile.Close()

	_, err = io.Copy(videoFile, videoStream)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully downloaded: %v", url)
	return nil
}

// // rsDownloadVideo downloads video using rust library
// func rsDownloadVideo(url string, path string) error {
// 	cUrl := C.CString(url)
// 	cPath := C.CString(path)
// 	C.download(cUrl, cPath)

// 	return nil
// }
