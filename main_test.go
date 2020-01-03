package main

import (
	"fmt"
	"testing"
)

func TestDownloader(t *testing.T) {
	urls := []string{"https://www.youtube.com/watch?v=HpNluHOAJFA&list=RDHpNluHOAJFA"}

	if err := beginDownload(urls); err != nil {
		fmt.Fprintln("Error occured testing downloader", err)
	}
}
