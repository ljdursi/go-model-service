# Build stage
FROM golang AS builder

WORKDIR /go/src/github.com/CanDIG/go-model-service
ADD . /go/src/github.com/CanDIG/go-model-service

# get sqlite3
RUN apt-get update && \
    apt-get install -y sqlite3 libsqlite3-dev

# get pop/soda with sqlite3 support compiled in
RUN go get -u -v -tags sqlite github.com/gobuffalo/pop/... && \
    go install -tags sqlite github.com/gobuffalo/pop/soda

# get dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# build go; CGO_ENABLED=0 and -a -installsuffix cgo for static;
# -tags sqlite to include sqlite3 in the binary
RUN cd /go/src/github.com/CanDIG/go-model-service && \
    dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags sqlite -o main ./variant-service/api/cmd/variant-service-server/main.go

# Final stage

FROM scratch

WORKDIR /
COPY --from=builder /go/src/github.com/CanDIG/go-model-service/main /

EXPOSE 3000

ENTRYPOINT ["/main", "--port=3000"]
