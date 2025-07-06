use axum::{Router, routing::get, response::Http};
use std::net::SocketAddr;

fn main() {
    let app = Router::new().route("/", get(route));

    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    println!("http://{}", addr);

}
