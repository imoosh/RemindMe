FROM golang:alpine as builder

WORKDIR /go/src/RemindMe
COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/RemindMe

COPY --from=0 /go/src/RemindMe ./

EXPOSE 8888

ENTRYPOINT ./server -c config.docker.yaml
