FROM golang:latest

RUN mkdir -p /go/src/github.com/Ch3ck/youtube-dl/

WORKDIR /go/src/github.com/Ch3ck/youtube-dl

COPY vendor     vendor
COPY api		api
COPY Makefile	Makefile
COPY ytd.go		.

RUN gofmt -l -d $(find . -type f -name '*.go' -not -path "./vendor/*") \
  && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o youtube-dl .


FROM alpine:latest

MAINTAINER Nyah Check <check.nyah@gmail.com>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	ca-certificates

WORKDIR /root/
COPY --from=0 /go/src/github.com/Ch3ck/youtube-dl .

RUN echo "Image build complete."


CMD [ "./youtube-dl" ]
