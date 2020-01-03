package main

import (
	"testing"
)

var tables = []struct {
	url, id string // input
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

	path := "test"
	for i, table := range tables {
		ID, _ := getVideoId(table.url)
		if ID != table.id {
			t.Errorf("videoId(%d): expected %q, actual %q", i, table.id, ID)
		}

		if ID != "" {
			if err := getVideoStream("mp3", ID, path, 192); err != nil {
				t.Errorf("videoStream(%d): expected %v, actual %v", i, nil, err)
			}
		}
	}
}

func TestVideoId(t *testing.T) {
	urls := []string{"https://www.youtube.com/watch?v=HpNluHOAJFA&list=RDHpNluHOAJFA"}

	url, err := getVideoId(urls[0])
	if err != nil {
		t.Log(err)
	}
	t.Log(url)
}

func BenchmarkVideoId(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getVideoId(tables[0].url)
	}
}

func BenchmarkApivideoStream(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getVideoStream("mp3", tables[0].id, "~/Downloads", 192)
	}
}

/*func BenchmarkApiConvertVideo(b *testing.B) {
	path := "~/Downloads/"
	for n := 0; n < b.N; n++ {
		file := path + tables[0].id + ".mp3"
		convertVideo(file, 123, tables[0].id, vid)
	}
}*/
