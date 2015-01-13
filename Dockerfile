FROM golang:latest
MAINTAINER Jess Frazelle <jess@docker.com>

RUN go get github.com/Sirupsen/logrus

ADD . /go/src/github.com/jfrazelle/battery
RUN cd /go/src/github.com/jfrazelle/battery && go install . ./...
ENV PATH $PATH:/go/bin

ENTRYPOINT ["battery"]
