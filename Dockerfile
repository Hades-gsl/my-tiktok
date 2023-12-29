# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

RUN mkdir /tiktok

WORKDIR /tiktok

COPY . .

RUN echo -e "https://mirrors.aliyun.com/alpine/v3.6/main/\nhttps://mirrors.aliyun.com/alpine/v3.6/community/" > /etc/apk/repositories \
    apk update && \
    apk add ffmpeg

RUN go mod tidy && \
    /tiktok/build.sh build

EXPOSE 8888

ENTRYPOINT sh build.sh run
