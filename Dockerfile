FROM golang:1.23-alpine

ENV GOOS=linux
ENV GOARCH=amd64

ARG VERSION=dev
ARG COMMIT=unknown

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -ldflags="-X 'main.version=${VERSION}' -X 'main.commit=${COMMIT}' -X 'main.date=$(date)'" -o pricing-engine-linux main.go
