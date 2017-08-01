# ytd
A Go Application for downloading Youtube videos and converting to other formats[FLV, mp3]. To detailed design architecture. Take a look at the [Design Document](ARCHITECTURE) markdown.


## Pre requisites

* [Go version 1.8](https://github.com/golang/go/releases/tag/go.1.8.3)
* [Docker CE 17.06](https://docs.docker.com/release-notes/docker-ce/)

Clone GIT repo:
```
$ git clone htps://github.com/Ch3ck/ytd
$ go get -u google.golang.org/api/youtube/v3
$ go get -u golang.org/x/oauth2/...
$ go get -u github.com/github.com/Sirupsen/logrus/...

```

## Build

The make command builds the code, runs the tests, generates and runs the docker containers.

```
$ make
```

## Kickstart usage

On a Linux or OSX system
```
$ ./ytd <link to youtube video>
```

## Roadmap

* Search for Youtube vidoes based on the Link and provides the download options for either mp3 or flv file
* Support HD Video download
* Support search with Youtube API, process results and user chooses whatever files to download
* Multithreaded downloads
* Web App(PWA, Basic JS Web UI)
* Package for OSX, Android, iOS


## Contributing

Start by starring and Forking this repository. Follow the basic instruction in the [CONTRIBUTING](CONTRIBUTING) file.

## Licence

YTD is licensed under [The MIT Licence](LICENSE).

## Author

This project was created and maintained by [Nyah Check](https://twitter.com/nyah_check). Please feel free to reach out, I could always use your help or advice :-)


## Notes

* Add vendoring for required Go deps.

