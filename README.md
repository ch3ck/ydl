# youtube-dl

[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/ch3ck/youtube-dl)
[![Github All Releases](https://img.shields.io/github/downloads/ch3ck/youtube-dl/total.svg?style=for-the-badge)](https://github.com/ch3ck/youtube-dl/releases)


`youtube-dl` is a simple youtube video downloader and can also download multiple videos concurrently.
Downloaded videos could be converted to `flv` or `mp3` formats.


## Dependencies

* [Go version 1.12](https://github.com/golang/go/releases/tag/go.1.12)
* [Docker CE 17.06](https://docs.docker.com/release-notes/docker-ce/)
* [Lame](https://sourceforge.net/projects/lame/)


## Build
```bash
$ git clone htps://github.com/Ch3ck/youtube-dl
$ cd youtube-dl
$ make
```

## Usage

```console
youtube-dl -h
youtube-dl - Simple youtube video/audio downloader

Usage: youtube-dl [OPTIONS] [ARGS]

Flags:
  -bitrate        Audio Bitrate (default 123)
  -format         File Format(mp3, webm, flv)
  -id             Youtube Video ID
  -path           Output Path (default ".")
  -version        print version and exit
  -h              Help page
```

## Roadmap

* Download youtube video with video id or link and converts to flv or mp3.
* Support HD Video download.
* Concurrent downloads.
* Web App(PWA, Basic JS Web UI).


## Contributing

Start by starring and Forking this repository. Follow the basic instruction in the [CONTRIBUTING](CONTRIBUTING.md) file.

## Licence

youtube-dl is licensed under [The MIT Licence](LICENSE.md).

## Author

This project was created and maintained by [Nyah Check](https://twitter.com/ch3ck_).
 Please feel free to reach out, I could always use your help or advice. :-)
