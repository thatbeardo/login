# start from golang image based on alpine-3.8
FROM golang:1.14-alpine3.11 AS dev-build

WORKDIR /go/src/app
ADD . /go/src/app

# Install User
RUN go get -d -v ./...
RUN go install -v ./...
# RUN make swagger

CMD ["sh", "-c", "$GOPATH/bin/userservice"]
