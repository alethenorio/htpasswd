#!/bin/bash -ex

NAME="htpasswd"
GOOS="linux"
GOARCH="amd64"

BINARY="dist/${NAME}-${GOOS}-${GOARCH}"

VERSION="$(git describe --tags)"

COMMITID="$(git rev-parse HEAD)"

BUILDTIME="$(date -u '+%Y-%m-%d_%H:%M:%S')"

echo "Building version=${VERSION} from commit=${COMMITID} for ${GOOS}/${GOARCH}"
CGO_ENABLED=0 go build -a -installsuffix cgo -o "${BINARY}" -ldflags "-X main.version=${VERSION} -X main.buildTime=${BUILDTIME} -X main.commitId=${COMMITID}"

GOOS=windows
GOARCH=amd64
BINARY="dist/${NAME}-${GOOS}-${GOARCH}.exe"

echo "Building version=${VERSION} from commit=${COMMITID} for ${GOOS}/${GOARCH}"
go build -o "${BINARY}" -ldflags "-X main.version=${VERSION} -X main.buildTime=${BUILDTIME} -X main.commitId=${COMMITID}"
