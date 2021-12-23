# ydl

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ch3ck/ydl/Build?style=for-the-badge)](https://github.com/ch3ck/ydl/actions)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/ch3ck/ydl)
[![GitHub license](https://img.shields.io/github/license/ch3ck/ydl?style=for-the-badge)](https://github.com/ch3ck/ydl/blob/master/LICENSE)

`ydl` is a simple youtube video downloader.


## Build

```bash
$ git clone https://github.com/Ch3ck/ydl
$ cd ydl
$ make
```


## Usage

To install ydl

```console
$ go get github.com/ch3ck/ydl
```

To run download:

```console
ydl -h
ydl - Simple youtube video/audio downloader

Usage: ydl [OPTIONS] [ARGS]

Flags:
  -bitrate        Audio Bitrate (default 123)
  -format         File Format(mp3, webm, flv)
  -id             Youtube Video ID
  -path           Output Path (default ".")
  -version        print version and exit
  -h              Help page
```

### Example

```bash
$ ./ydl -format mp3 -id lWEbEtr_Vng
```

## Roadmap

* Download youtube video with video id or link and converts to flv or mp3.
* Support HD Video download.
* Concurrent downloads.
* Web App(PWA, Basic JS Web UI).


## Contributing

Follow the basic instruction in the [CONTRIBUTING](CONTRIBUTING.md) file.

## Licence

`ydl` is licensed under [The MIT Licence](LICENSE.md).


## Support

This project was created and is maintained by [Nyah Check](https://twitter.com/ch3ck_)
