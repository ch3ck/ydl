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

func TestApi(t *testing.T) {

	path := "~/Downloads/"
	for i, table := range tables {
		var rawVideo *RawVideoData
		ID, _ := GetVideoId(table.url)
		if ID != table.id {
			t.Errorf("GetVideoId(%d): expected %q, actual %q", i, table.id, ID)
		}

		if ID != "" {
			video, err := APIGetVideoStream(ID, rawVideo)
			if err != nil {
				t.Errorf("APIGetVideoStream(%d): expected %v, actual %v", i, nil, err)
			}

			file := path + table.url + ".mp3"
			err = APIConvertVideo(file, 123, ID, video)
			if err != nil {
				t.Errorf("APIConvertVideo(%d): expected %v, actual %v", i, nil, err)
			}
		}
	}
}
