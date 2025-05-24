# Build
FROM    golang:latest AS build

ENV     GO111MODULE=on
ENV     CGO_ENABLED=0

RUN     mkdir /go/src/app
WORKDIR /go/src/app
ADD     . /go/src/app
ENV 	GOCACHE=/root/.cache/go-build
RUN 	--mount=type=cache,target="/root/.cache/go-build"
RUN     go get github.com/labstack/echo/v5@v5.0.0-20230722203903-ec5b858dab61
#RUN     make deps
RUN     make build
EXPOSE  8080 
EXPOSE  4000