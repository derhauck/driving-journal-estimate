FROM golang:1.27.4 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ cmd/
COPY factory/ factory/
COPY internal/ internal/
COPY public/ public/
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags='-s -w' -o /out/dje .

# Static binary: no shell or libc needed. The pod's preStop must use the
# native sleep handler, not `exec sleep`.
FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=builder /out/dje /usr/bin/dje
COPY schema/oas.yml /app/schema/openapi.yml
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/usr/bin/dje"]
CMD ["calendar", "server"]
