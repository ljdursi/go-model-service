# Build stage
FROM golang
ARG WORKDIR=/go/src/github.com/CanDIG/go-model-service

ADD . /go/src/github.com/CanDIG/go-model-service

# Install CLI tool dependencies for install.sh script
RUN cd $WORKDIR ./install_dep.sh

# build go
# -tags sqlite to include sqlite3 in the binary
RUN cd $WORKDIR && ./install.sh

EXPOSE 3000

ENTRYPOINT ["/go/src/github.com/CanDIG/go-model-service/main", "--port=3000"]
