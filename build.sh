NAME="htpasswd"
BINARY="dist/${NAME}-${GOOS}-${GOARCH}"
VERSION="${TRAVIS_TAG}"
COMMIT="${TRAVIS_COMMIT}"

if [ -z "$VERSION" ]; then
    VERSION="${COMMIT}"
fi

if [ "$GOOS" == "windows" ]; then
    BINARY="${BINARY}.exe"
fi

tar cfvz "${NAME}-${VERSION}-tar.gz" .

go build -o "${BINARY}" -ldflags "-X main.Version=${VERSION} -X main.buildTime=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.commitId=${TRAVIS_COMMIT}"

mv "${NAME}-${VERSION}-tar.gz" dist/