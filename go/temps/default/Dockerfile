# syntax=docker/dockerfile:1

########################
# build stage
########################
FROM golang:1.24-alpine AS builder
WORKDIR /src

# cache dependencies
COPY go.mod .
# COPY go.sum .   # uncomment if you have go.sum
RUN go mod download

# copy source & compile
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./...

########################
# runtime stage
########################
FROM alpine:latest
WORKDIR /app
COPY --from=builder /src/app .
ENTRYPOINT ["./app"]

