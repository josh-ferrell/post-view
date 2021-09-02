FROM debian:latest

ADD post-view /

ENTRYPOINT ["/post-view"]
