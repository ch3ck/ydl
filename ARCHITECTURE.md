# ARCHITECTURE

YTD is a tool for searching and downloading video and audio content from Youtube. It includes primitives for basic data retrieval, content search, download and conversion to other audio and video file formats. 


## Overview

It contains three main packages namely: 
* the API(Data retrieval and requests to the Youtube API), 
* Auth(Uses OAuth to make authenticated requests by the API package) when connecting to the Youtube API. There also consists a converter which is part of the API package to convert downloaded youtube data into other file formats such as FLV and mp3 only. 
* The CMD basically gets the tool up and running taking queries from the user and returning data output to the user.

  ----------------------          =========================
  |
  |						  OAuth		
  |   API               |  -----> |	   Youtube API
  |						  
  |
  ----------------------          =========================

		^
		'
		'
		'
		
		
   ---------------------					-------------------
   					   |					|				  |
   		Auth				<------->              CMD
   					   |				    |                 |
   ---------------------                    -------------------

	--> = Channel

	
## Features

The basic features are listed here below

* YTD API
 **- apidata.go:	Sends authenticated requests to the Google Youtube API
 					Process video feeds and downloads video data from Youtube.
 **- apisearch.go:	Performs search queries on Youtube and parses the results to `apidata.go`
 **- apiconv.go:	Converts downloaded videos to FLV and mp3
 
* YTD Auth
 **- auth.go:	Uses OAuth to perform authenticated requests with Youtube, creating the `YoutubeRequestSettings` and `YoutubeSettings` objects.

* YTD CMD
 **- main.go:	Program entry point, gets queries from the user.
 				Creates requests via `AUTH` API and connects to `API` package and displays results to user.

* YTD UI
 **- ui.go:		UI creates GUI interface for parsing of response data to user(UI/UX)


## Build

Requirements:

* Go 1.8 or higher
* Docker CE 17
* A working Golang environment
* Protobuf 3.x or higher to regenerate protocol buffer files (e.g. using make generate)

`ytd` is built in Go using the standard `go` project structure to work well with Go tooling.


## Test

Tests are ran by running `make` or `make all`

```
$ make
```
---
