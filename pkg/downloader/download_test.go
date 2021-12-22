package downloader

import (
	"testing"
)

var tables = []struct {
	url, id string // input
}{
	{"https://www.youtube.com/watch?v=lWEbEtr_Vng", "lWEbEtr_Vng"},
	{"https://www.youtube.com/watch?v=ALWmcO8S-dc", "ALWmcO8S-dc"},
	{"", ""},
	{"https://www.facebook.com/mark/videos?v=RDHpNluHOAJFA", ""},
	{"https://www.youtube.com/watch?v=ALWmcO8S-dc", "ALWmcO8S-dc"},
	{"https://www.wsj.com/articles/trump-administration-wont-withdraw-from-paris-climate-deal-1505593922", ""},
	{"https://vimeo.com/101522071", ""},
}

var vid []string

func TestApi(t *testing.T) {

	// path := "test"
	for i, table := range tables {
		err := DecodeVideoStream(table.url, "mp3")
		if err != nil {
			t.Errorf("videoId(%d): expected %q, actual %q", i, table.id, err)
		}
	}
}

func BenchmarkVideoId(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if err := DecodeVideoStream(tables[0].url, "mp3"); err != nil {
			b.Errorf("Error downloading video: %v", err)
		}
	}
}
