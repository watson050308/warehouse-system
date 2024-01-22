FROM --platform=linux/amd64 golang:1.20 AS builder

WORKDIR /build

COPY . .
RUN go mod download

FROM builder AS build 

ENV GOMODCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" \
    GOOS=linux GOARCH=amd64 go build -o /app .

ENTRYPOINT ["/app"]
