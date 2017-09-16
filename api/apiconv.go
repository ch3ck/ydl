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
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Sirupsen/logrus"
)

//Converts Decoded Video file to mp3 by default with 123 bitrate or to
//flv if otherwise specified and downloads to system
func APIConvertVideo(file string, bitrate uint, id string, decVideo []string) error {
	cmd := exec.Command("ffmpeg", "-i", "-", "-ab", fmt.Sprintf("%dk", bitrate), file)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if filepath.Ext(file) != ".mp3" && filepath.Ext(file) != ".flv" {
		file = file[:len(file)-4] + ".mp3"
	}

	logrus.Infof("Converting video to %q format", filepath.Ext(file))
	if filepath.Ext(file) == ".mp3" {
		/* NOTE: To modify to use Go ffmpeg bindings or cgo */

		buf := &bytes.Buffer{}
		gob.NewEncoder(buf).Encode(decVideo)
		_, err = exec.LookPath("ffmpeg")
		if err != nil {
			return errors.New("ffmpeg not found on system")
		}

		cmd.Start()
		logrus.Infof("Downloading mp3 file to disk %s", file)
		stdin.Write(buf.Bytes()) //download file.

	} else {
		out, err := os.Create(file)
		if err != nil {
			logrus.Errorf("Unable to download video file.", err)
			return err
		}
		err = apiDownloadVideo(id, out)
		return err
	}

	return nil
}

//Downloads decoded video stream.
func apiDownloadVideo(url string, out io.Writer) error {
	logrus.Infof("Downloading file stream")

	resp, err := http.Get(videoExtractor + url)
	if err != nil {
		return fmt.Errorf("requesting stream: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("reading answer: non 200 status code received: '%s'", err)
	}
	length, err := io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("saving file: %s (%d bytes copied)", err, length)
	}

	logrus.Infof("Downloaded %d bytes", length)

	return nil
}
