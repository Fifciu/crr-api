FROM golang:1.11
LABEL maintainer="Filip Jędrasik <filip.jdrasik@gmail.com>"
WORKDIR /go/src/github.com/filipjedrasik/crr-api

COPY . .
COPY ./wait-for-it.sh ./

RUN go get -d -v ./...; exit 0
RUN go install -v ./...
CMD ["go", "run", "main.go"]
