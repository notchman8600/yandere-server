ARG SHA256_KEY="hogehoge"

FROM golang:latest as builder
WORKDIR /workdir
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o app

# Runtime Container
FROM alpine
ARG SHA256_KEY
ENV SHA256_KEY=$SHA256_KEY
RUN apk add --no-cache ca-certificates
COPY --from=builder /workdir/app /app

ENTRYPOINT ["/app"]
