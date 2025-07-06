use axum::{Router, routing::get, response::Html};
use std::net::SocketAddr;


#[tokio::main]
async fn main() {
    let app = Router::new().route("/" get(route));

    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    println!("http://{addr}", addr);

    axum::Server::bind(&addr)
        .serve(app.into_make_serive())
        .await
        .unwrap();
}

