#[cfg(test)]
mod main_test;

fn main() {
    let nums = vec![1, 2, 3, 4];
    let target = 7;
    let result = run(&nums, target);

    println!("Result: {:?}", result)
}
