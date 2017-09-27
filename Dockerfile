FROM golang:latest

RUN mkdir -p /go/src/github.com/Ch3ck/youtube-dl/

WORKDIR /go/src/github.com/Ch3ck/youtube-dl

COPY vendor     vendor
COPY ytd.go  .

RUN gofmt -l -d $(find . -type f -name '*.go' -not -path "./vendor/*") \
  && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ytd .


FROM alpine:latest

MAINTAINER Nyah Check <check.nyah@gmail.com>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache add ca-certificates

WORKDIR /root/
COPY --from=0 /go/src/github.com/Ch3ck/ytd .

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		go \
		git \
		gcc \
		libc-dev \
		libgcc \
		ffmpeg \
	&& cd /go/src/github.com/Ch3ck/ytd \
	&& go build -o /usr/bin/ytd . \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."


CMD [ "./ytd" ]
