FROM alpine:latest

ADD post-view /

ENTRYPOINT ["/post-view"]
