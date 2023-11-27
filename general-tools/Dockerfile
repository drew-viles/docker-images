FROM cgr.dev/chainguard/wolfi-base:latest

LABEL "uk.hudson-viles.registry"="General Tools Image"
LABEL uk.hudson-viles.registry-image-base="wolfi-base:latest"
LABEL description="This image can be used as a base for all other images and contains generic tools."

RUN apk update
RUN apk --no-cache add curl wget jq bash ca-certificates openssl bind-tools
