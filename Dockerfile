FROM sconecuratedimages/crosscompilers
COPY / /
RUN apk update \
    && apk add git curl go \
    && export SCONE_HEAP=1G \
    && go get ./... \
    && go build -compiler gccgo -buildmode=exe -o app cmd/v0/main.go

RUN ldd groupcache

FROM alpine:latest
CMD sh -c "SCONE_HEAP=1G /app"