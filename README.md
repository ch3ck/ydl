# ytd
`ytd` is Go application for downloading Youtube videos and converting to other formats(flv, mp3). To understand the design take a look at the [Design Document](ARCHITECTURE.md).

```
.
├── api
│   ├── apiconv.go
│   ├── apidata.go
│   └── apidata_test.go
├── ARCHITECTURE.md
├── build.sh
├── cmd
│   └── ytd
│       ├── ytd.go
│       └── ytd_test.go
├── CONTRIBUTING.md
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── simple.go
└── vendor
      ....
```

## Pre requisites

* [Go version 1.8](https://github.com/golang/go/releases/tag/go.1.8.3)
* [Docker CE 17.06](https://docs.docker.com/release-notes/docker-ce/)
* [Lame](https://sourceforge.net/projects/lame/)

Clone GIT repo and install some dependencies:
```
$ git clone https://github.com/Ch3ck/youtube-dl

```

## Build

The make command builds the code, runs the tests, generates and runs the docker containers.

```
$ make
```

## Kickstart usage

On a Linux or OSX system
```
ytd -id 'videoId' -format mp3 -bitrate 123  -path ~/Downloads/ videoUrlv0.1
  -bitrate uint
    	Audio Bitrate (default 123)
  -format string
    	File Format(mp3, webm, flv)
  -id string
    	Youtube Video ID
  -path string
    	Output Path (default ".")
  -version
    	print version and exit

  searches for youtube video on that link and downloads
```
Running `ytd` without any arguments will prompt for link.


## Roadmap

* Search for Youtube vidoes based on the Link and provides the download options for either mp3 or flv file.
* Support HD Video download.
* Multithreaded downloads.
* Web App(PWA, Basic JS Web UI).
* Package for OSX, Android, iOS.


## Contributing

Start by starring and Forking this repository. Follow the basic instruction in the [CONTRIBUTING](CONTRIBUTING.md) file.

## Licence

YTD is licensed under [The MIT Licence](LICENSE.md).

## Author

This project was created and maintained by [Nyah Check](https://twitter.com/nyah_check). Please feel free to reach out, I could always use your help or advice. :-)
