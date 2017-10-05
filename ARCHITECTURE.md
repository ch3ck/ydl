# ARCHITECTURE

YTD is a tool for searching and downloading video and audio content from Youtube. It includes primitives for basic data retrieval, content search, download and conversion to other audio and video file formats. 



## Overview

It contains three main packages namely: 
* API: Gets video Id, Gets Video Information, Decodes Video info, Convert to appropriate stream and Downloads
* CMD: Main package that starts the program entry point.
* UI:  Creates A UI for getting user input and returning results.

	
## Features

The basic features are listed here below

* YTD API
 **- apidata.go: Get Video ID, Get Video Info, Decode Video Info, Convert to proper Stream, Download Information

* YTD CMD
 **- main.go:	Program entry point, gets queries from the user.

* YTD UI
 **- ui.go:		UI creates GUI interface for parsing of response data to user(UI/UX)


## Build

Requirements:

* Go 1.8 or higher
* Docker CE 17
* A working Golang environment

`ytd` is built in Go using the standard `go` project structure to work well with Go tooling.


## Test

Tests are run using `make` or `make all`

```
$ make
```
---
