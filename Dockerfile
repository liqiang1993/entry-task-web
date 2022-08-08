FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/phoenix_gateway
COPY . $GOPATH/src/phoenix_gateway
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./phoenix_gateway"]
