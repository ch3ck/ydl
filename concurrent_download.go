/**
 * download video files concurrently
 */
package main

import "sync"

//DownloadStreams download a batch of elements asynchronously
func DownloadStreams(maxOperations int, format, outputPath string, bitrate uint, urls []string) <-chan error {

	var wg sync.WaitGroup
	wg.Add(len(urls))

	ch := make(chan error, maxOperations)
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()

			if ID, err := getVideoId(url); err != nil {
				ch <- err
			} else {
				ch <- getVideoStream(format, ID, outputPath, bitrate)
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
