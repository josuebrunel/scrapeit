# Build
FROM    golang:latest AS build

ENV     GO111MODULE=on
ENV     CGO_ENABLED=0

RUN     mkdir /go/src/app
WORKDIR /go/src/app
ADD     . /go/src/app
ENV 	GOCACHE=/root/.cache/go-build
RUN 	--mount=type=cache,target="/root/.cache/go-build"
RUN     make deps
# RUN     make build
EXPOSE  8080 
EXPOSE  7654

# Publisher
FROM alpine:latest AS publisher
RUN  mkdir /opt/scrapeit/
WORKDIR /opt/scrapeit/
COPY --from=build  /go/src/app/bin/scrapeit-server /opt/scrapeit/scrapeit-server
COPY    --from=build /go/src/app/static/ /opt/scrapeit/static/

EXPOSE 8080
EXPOSE 7654
CMD [ "./scrapeit-server" ]

# Subscriber
FROM alpine:latest AS subscriber
RUN  mkdir /opt/scrapeit/
WORKDIR /opt/scrapeit/
COPY --from=build  /go/src/app/bin/scrapeit-client /opt/scrapeit/scrapeit-client
CMD [ "./scrapeit-client" ]
