FROM gcr.io/oss-fuzz-base/base-builder-go:v1
COPY . $GOPATH/src/google.golang.org/protobuf
WORKDIR $GOPATH/src/google.golang.org/protobuf
COPY .clusterfuzzlite/build.sh $SRC/
