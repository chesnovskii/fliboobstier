FROM golang:alpine3.8 as build
COPY . /go/src/github.com/chesnovsky/fliboobstier
WORKDIR /go/src/github.com/chesnovsky/fliboobstier
RUN  apk add --no-cache git make gcc libc-dev ca-certificates \
  && make deps \
  && make

FROM library/alpine:3.8
RUN apk add --no-cache ca-certificates
COPY config.yml /config.yml
COPY --from=build /go/src/github.com/chesnovsky/fliboobstier/bin /fliboobstier
CMD ["/fliboobstier"]
