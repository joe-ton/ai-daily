use std::collections::HashMap;

fn main() {
    let nums = vec![1, 2, 3, 4];
    let target = 7;
    let result = run(&nums, target); 

    println!("Hello, {:?}!", result);
}

fn run(nums: &[i32], target: i32) -> Vec<usize> {
    if nums.len() < 2 {
        return vec![]
    }

    let mut seen = HashMap::new();

    for (idx, num) in nums.iter().enumerate() {
        let complement = target - num;

        if let Some(&comp_idx) = seen.get(&complement) {
            return vec![comp_idx, idx];
        }
        seen.insert(num, idx);
    }
    vec![]
}

