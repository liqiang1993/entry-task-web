FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/web-server
COPY . $GOPATH/src/web-server
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./web-server"]
