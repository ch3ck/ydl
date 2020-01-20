
# Build container
FROM golang:1.13-alpine AS go-base
RUN apk add --no-cache git

MAINTAINER Nyah Check <hello@nyah.dev>

ENV CGO_ENABLED=1
ENV GO111MODULE=on
WORKDIR /app
RUN echo "Build container"
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /youtube-dl .


# Runtime container
FROM scratch
RUN echo "Runtime container"
COPY --from=go-base /app/youtube-dl /youtube-dl
ENTRYPOINT ["/youtube-dl"]
