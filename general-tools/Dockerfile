FROM alpine:3.16

LABEL "com.viles.registry"="General Tools Image"
LABEL com.viles.registry-image-base="alpine:3.16.1"
LABEL version="0.0.1"
LABEL description="This image can be used as a base for all other images and contains generic tools."

RUN apk update
RUN apk --no-cache add curl wget jq bash ca-certificates openssl bind-tools