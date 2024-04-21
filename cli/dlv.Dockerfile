FROM golang:1.21.5
WORKDIR /go/src/
RUN go install github.com/go-delve/delve/cmd/dlv@latest
CMD ["/go/bin/dlv", "debug", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient"]
EXPOSE 40000