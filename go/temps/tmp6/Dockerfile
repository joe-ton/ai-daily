# build stage

FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

# runtime stage
FROM alpine:latest
WORKDIR /

COPY --from=builder /app .
ENTRYPOINT ["./app"]


