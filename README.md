# ytd
`ytd` is Go application for downloading Youtube videos and converting to other formats(flv, mp3). To understand the design take a look at the  [Design Document](ARCHITECTURE.md).

```
ytd
├── api
│   ├── apiconv.go
│   ├── apiconv_test.go
│   ├── apidata.go
│   ├── apidata_test.go
│   ├── apisearch.go
│   └── apisearch_test.go
├── ARCHITECTURE.md
├── auth
│   ├── auth.go
│   ├── auth_test.go
│   ├── client_secret.json
│   └── ytd-auth.json
├── cmd
│   └── ytd
│       ├── ytd.go
│       └── ytd_test.go
├── CONTRIBUTING.md
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
└── vendor

```

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
$ ytd -h

  -s or --search 
        searches for any matching videos on Youtube
  -k or --key
        Youtube API token
  -v or --version    print version and exit (shorthand)
		prints ytd version and exits
		
$ ytd <link>
  searches for youtube video on that link and downloads
```
Running `ytd` without any arguments will prompt for link


## Roadmap

* Search for Youtube vidoes based on the Link and provides the download options for either mp3 or flv file
* Support HD Video download
* Support search with Youtube API, process results and user chooses whatever files to download
* Multithreaded downloads
* Web App(PWA, Basic JS Web UI)
* Package for OSX, Android, iOS


## Contributing

Start by starring and Forking this repository. Follow the basic instruction in the [CONTRIBUTING](CONTRIBUTING.md) file.

## Licence

YTD is licensed under [The MIT Licence](LICENSE.md).

## Author

This project was created and maintained by [Nyah Check](https://twitter.com/nyah_check). Please feel free to reach out, I could always use your help or advice :-)
