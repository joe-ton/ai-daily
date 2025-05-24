use std::collections::HashMap();

fn main() {
    println!("Hello, world!");
}

pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut seen = HashMap::new();

    for (idx, num) in nums.iter().enumerate() {
        let complement = target - num;
        if let Some(&comp_idx) = seen.get(&complement) {

        }
    }
}
