FROM golang:1.24-alpine AS builder
WORKDIR /cmd

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOUS=linux go build -o app ./...

FROM alpine:latest
WORKDIR /app
COPY --from=builder /cmd/app .
ENTRYPOINT ["./app"]
