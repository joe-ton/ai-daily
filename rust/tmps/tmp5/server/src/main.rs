use axum::{Router, routing::get, response::Html};
use std::net::SocketAddr;

#[tokio::main]
async fn main() {
    let app = Router::new().get(route);

    let addr = SockerAddr::from(([127, 0, 0, 1], 3000));
    println!("Http://{addr}", addr);
}

async fn root() -> 
