FROM sconecuratedimages/crosscompilers
COPY ./ /go/src/github.com/carapace/core
RUN apk update \
    && apk add git curl \
    && wget https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh \
    && bash goinstall.sh --64 \
    && source /root/. \
    && export SCONE_HEAP=1G \
    && export GOPATH=$HOME/go \
    && go get -d -v ./... \
    && go build -compiler gccgo -buildmode=exe -gccgoflags -g /go/src/github.com/carapace/core/cmd/v0/main.go

RUN ldd groupcache

FROM alpine:latest
CMD sh -c "SCONE_HEAP=1G /app"