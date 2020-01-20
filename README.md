# youtube-dl

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ch3ck/youtube-dl/Build?style=for-the-badge)](https://github.com/ch3ck/youtube-dl/actions)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/ch3ck/youtube-dl)
[![GitHub license](https://img.shields.io/github/license/ch3ck/youtube-dl?style=for-the-badge)](https://github.com/ch3ck/youtube-dl/blob/master/LICENSE)

`youtube-dl` is a simple youtube video downloader and can also download multiple videos concurrently.
Downloaded videos could be converted to `flv` or `mp3` formats.


## Build

```bash
$ git clone htps://github.com/Ch3ck/youtube-dl
$ cd youtube-dl
$ make
```

## Install



## Usage

To install youtube-dl

```console
$ go get github.com/Ch3ck/youtube-dl
```

To run download:

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

### Example

```console
$ ./youtube-dl -format mp3 -id https://www.youtube.com/watch?v=lWEbEtr_Vng
```

## Roadmap

* Download youtube video with video id or link and converts to flv or mp3.
* Support HD Video download.
* Concurrent downloads.
* Web App(PWA, Basic JS Web UI).


## Contributing

Follow the basic instruction in the [CONTRIBUTING](CONTRIBUTING.md) file.

## Licence

`youtube-dl` is licensed under [The MIT Licence](LICENSE.md).


## Support

This project was created and is maintained by [Nyah Check](https://twitter.com/ch3ck_)