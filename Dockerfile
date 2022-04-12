FROM golang:alpine as builder

WORKDIR /go/src/github.com/flipped-aurora/gin-vue-admin/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env -w GOOS=linux \
    && go env \
    && go mod tidy \
    && go build -ldflags="-w -s" -a -installsuffix cgo -o main .

FROM alpine:latest

LABEL MAINTAINER="re_chinese@re.com"

WORKDIR /go/release

COPY --from=builder /go/release/server ./

EXPOSE 8000
CMD ["/server","server","-c"]
