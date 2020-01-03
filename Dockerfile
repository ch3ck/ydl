FROM golang:latest AS go-base

MAINTAINER Nyah Check <hello@nyah.dev>

RUN apk add --no-cache \
	ca-certificates \
	make

FROM go-base
WORKDIR /go/src/github.com/ch3ck/youtube-dl
COPY . .
RUN make build

RUN echo "Image build complete."


CMD [ "./youtube-dl" ]