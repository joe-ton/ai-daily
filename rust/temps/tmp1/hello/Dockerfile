FROM rustlang/rust:nightly AS builder
WORKDIR /app

COPY . .
RUN cargo build --release

FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/target/release/hello .
CMD ["./hello"]
