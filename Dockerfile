FROM golang:alpine as builder

WORKDIR /go/release/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env -w GOOS=linux \
    && go env \
    && go mod tidy \
    && go build -ldflags="-w -s" -a -installsuffix cgo -o main main.go

FROM alpine:latest

LABEL MAINTAINER="re_chinese@re.com"

COPY --from=builder /go/release/server ./

EXPOSE 8080
CMD ["/main","main","-c"]