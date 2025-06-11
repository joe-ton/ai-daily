# build stage
FROM rust:1.85-alpine AS builder
WORKDIR /hello

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOUS=linux go build -o app ./...

#runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /src/app .
ENTRYPOINT ["./app"]
