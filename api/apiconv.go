// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// apiconv: Converts Decoded Video data to MP3, WEBM or MP4.
// NOTE: To reimplement using Go ffmpeg bindings.
package api

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

//Converts Decoded Video file to mp3 by default with 123 bitrate or to
//flv if otherwise specified and downloads to system
func ApiConvertVideo(file, id, format string, bitrate uint, decVideo []string) error {
	cmd := exec.Command("ffmpeg", "-i", "-", "-ab", fmt.Sprintf("%dk", bitrate), file)
	/* if err := os.MkdirAll(filepath.Dir(file), 666); err != nil {
		return err
	}
	out, err := os.Create(file)
	if err != nil {
		return err
	} */

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(decVideo)
	_, err = exec.LookPath("ffmpeg")
	if err != nil {
		return errors.New("ffmpeg not found on system")
	}

	cmd.Start()
	//logrus.Infof("Downloading mp3 file to disk %s", file)
	stdin.Write(buf.Bytes()) //download file.

	return nil
}

//Downloads decoded video stream.
func ApiDownloadVideo(file string, url string, video *RawVideoData) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Http.Get\nerror: %s\nURL: %s\n", err, url)
		return err
	}
	defer resp.Body.Close()
	video.Vlength = float64(resp.ContentLength)

	if resp.StatusCode != 200 {
		log.Printf("Reading Output: status code : '%v'", resp.StatusCode)
		return errors.New("non 200 status code received")
	}
	err = os.MkdirAll(filepath.Dir(file), 666)
	if err != nil {
		return err
	}
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	mw := io.MultiWriter(out, video)
	_, err = io.Copy(mw, resp.Body)
	if err != nil {
		log.Println("Download Error: ", err)
		return err
	}
	return nil
}
