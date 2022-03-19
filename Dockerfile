FROM golang:alpine
RUN apk add git

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest


COPY . .
