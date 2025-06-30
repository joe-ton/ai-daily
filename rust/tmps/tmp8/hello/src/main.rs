


fn main() {
    let task = env::args().nth(1).expect("usage: todoer <task text>");
    println("Added: {}", task);
}
