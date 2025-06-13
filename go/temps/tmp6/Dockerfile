
# build stage
FROM golang:1.24-alpine AS builder
WORKDIR /cmd

COPY go.mod .
RUN go mod download


COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

# runtime stage
FROM alpine:latests
WORKDIR /app
COPY --from=builder /cmd/app .
ENTRYPOINT ["./app"]
