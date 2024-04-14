FROM golang:1.22 as builder
RUN mkdir -p /app
WORKDIR /app
COPY cmd/ /app/cmd
COPY factory/ /app/factory
COPY internal/ /app/internal
COPY public/ /app/public
COPY main.go go.mod go.sum /app/
#COPY . /app/
ENV GOMODCACHE /app/.pkg
ENV GOCACHE /app/.cache


RUN CGO_ENABLED=0 GOOS=linux go build -o . -v ./...


FROM alpine:3.19

COPY --from=builder /app/driving-journal-estimate /usr/bin/dje
EXPOSE 8080
CMD dje calendar server
