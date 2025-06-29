# builder stage
FROM rust:1.79-bookworm AS builder
WORKDIR /app

COPY Cargo.toml Cargo.lock ./
RUN mkdir src && echo "fn main(){}" > src/main.rs
RUN cargo fetch --locked
RUN cargo build --release

COPY src ./src
RUN cargo build --release --bin hello

# runtime stage
FROM gcr.io/distroless/cc-debian12 AS runtime
LABEL org.opencontainers.image.source="https://github.com/your/repo"

COPY --from=builder /app/target/release/hello /usr/local/bin/hello
USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/hello"]
