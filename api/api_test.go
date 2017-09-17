// Copyright 2017 YTD Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// api_test: Tests API package.

package api

import (
	"testing"
)

var tables = []struct {
	url string // input
	id  string // expected result
}{
	{"https://www.youtube.com/watch?v=HpNluHOAJFA&list=RDHpNluHOAJFA", "HpNluHOAJFA"},
	{"https://www.youtube.com/watch?v=jOWsu8ePrbE&list=RDHpNluHOAJFA&index=8", "jOWsu8ePrbE"},
	{"", ""},
	{"https://www.facebook.com/mark/videos?v=RDHpNluHOAJFA", ""},
	{"https://www.youtube.com/watch?v=lWEbEtr_Vng", "lWEbEtr_Vng"},
	{"https://www.wsj.com/articles/trump-administration-wont-withdraw-from-paris-climate-deal-1505593922", ""},
	{"https://vimeo.com/101522071", ""},
}

var vid []string

func TestApi(t *testing.T) {

	//path := "~/Downloads/"
	for i, table := range tables {
		var rawVideo *RawVideoData
		ID, _ := GetVideoId(table.url)
		if ID != table.id {
			t.Errorf("GetVideoId(%d): expected %q, actual %q", i, table.id, ID)
		}

		if ID != "" {
			_ , err := APIGetVideoStream(ID, rawVideo)
			if err != nil {
				t.Errorf("APIGetVideoStream(%d): expected %v, actual %v", i, nil, err)
			}

			/*file := path + table.id + ".mp3"
			err = APIConvertVideo(file, 123, ID, video)
			if err != nil {
				t.Errorf("APIConvertVideo(%d): expected %v, actual %v", i, nil, err)
			}*/
		}
	}
}

func BenchmarkGetVideoId(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetVideoId(tables[0].url)
	}
}

func BenchmarkApiGetVideoStream(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var rawVideo *RawVideoData
		vid, _ = APIGetVideoStream(tables[0].id, rawVideo)
	}
}
/*
func benchmarkApiConvertVideo(b *testing.B) {
	path := "~/Downloads/"
	for n := 0; n < b.N; n++ {
		file := path + tables[0].id + ".mp3"
		APIConvertVideo(file, 123,tables[0].id, vid)
	}
} */
