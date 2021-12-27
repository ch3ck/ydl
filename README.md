# ydl

[![Build](https://github.com/ch3ck/youtube-dl/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/ch3ck/youtube-dl/actions/workflows/ci.yml)
[![release new binaries](https://github.com/ch3ck/ydl/actions/workflows/release.yml/badge.svg?branch=v0.1.0)](https://github.com/ch3ck/ydl/actions/workflows/release.yml)
[![CodeQL](https://github.com/ch3ck/youtube-dl/actions/workflows/codeql-analysis.yml/badge.svg?branch=master)](https://github.com/ch3ck/youtube-dl/actions/workflows/codeql-analysis.yml))
[![forthebadge](https://forthebadge.com/images/badges/contains-technical-debt.svg)](https://forthebadge.com)

`ydl` is a simple youtube video downloader.


## Build

### Pre-requisites

1. Install [rust nightly](https://rust-lang.github.io/rustup/concepts/channels.html)
2. Install the [go](https://go.dev/doc/install)


### Install and Run
![build and run instructions](https://user-images.githubusercontent.com/4006891/147508842-fdc95517-995e-494f-ad8a-f2860eb7a0ce.png)


### Rust bindings

As for the rust package, you can use cgo to compile and rust, but you'll have to add the following lines to the `main.go`

![building with rust bindings](https://user-images.githubusercontent.com/4006891/147509114-d49f733f-db91-459e-b860-6f932f19e8d0.png)

then run
```bash
$ make build-static
```


## License
The scripts and documentation in this project are released under the [MIT License](LICENSE.md)


## Author

**NOTE** Will stop maintaining this package soon, the *Youtube API* is highly unreliable(its a rat race).

- [Nyah Check](https://nyah.dev)
