// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// apiconv: Converts Decoded Video data to MP3, WEBM or MP4.
// NOTE: To reimplement using Go ffmpeg bindings.
package api

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/viert/lame"
)

//Downloads decoded audio stream
func ApiConvertVideo(file, path string, bitrate uint, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Http.Get\nerror: %s\nURL: %s\n", err, url)
		return err
	}
	defer resp.Body.Close()

	data, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		logrus.Errorf("Error reading video data: %v", e)
	}

	curDir, er := user.Current()
	if er != nil {
		return er
	}

	homeDir := curDir.HomeDir
	dir := homeDir + "/Downloads/youtube-dl/" + path
	fp := filepath.Join(dir, file)
	if err := os.MkdirAll(filepath.Dir(fp), 0775); err != nil {
		return err
	}

	os.Remove(fp) //delete if file exists.
	out, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer out.Close()
	r := bytes.NewReader(data)
	reader := bufio.NewReader(r)
	audioWriter := lame.NewWriter(out)
	audioWriter.Encoder.SetBitrate(int(bitrate))
	audioWriter.Encoder.SetQuality(1)

	// IMPORTANT!
	audioWriter.Encoder.InitParams()
	reader.WriteTo(audioWriter)

	return nil
}

//Downloads decoded video stream.
func ApiDownloadVideo(path, file, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Http.Get\nerror: %s\nURL: %s\n", err, url)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Reading Output: status code: '%v'", resp.StatusCode)
		return errors.New("Non 200 status code received")
	}

	curDir, er := user.Current()
	if er != nil {
		return er
	}
	homeDir := curDir.HomeDir
	dir := homeDir + "/Downloads/youtube-dl/" + path
	fp := filepath.Join(dir, file)
	err = os.MkdirAll(filepath.Dir(fp), 0775)
	if err != nil {
		return err
	}
	os.Remove(fp) //delete if file exists
	out, err := os.Create(fp)
	if err != nil {
		return err
	}

	//saving downloaded file.
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Download Error: ", err)
		return err
	}
	return nil
}
