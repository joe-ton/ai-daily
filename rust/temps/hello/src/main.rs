use std::collections::HashMap;

fn main() {
    let nums = vec![1, 2, 3, 4];
    let target = 7;
    let result = two_sum(&nums, target);

    println!("Result: {:?}", result);
}

pub fn two_sum(nums: &[i32], target: i32) -> Vec<usize> {
    let mut seen = HashMap::new();

    for (idx, num) in nums.iter().enumerate() {
        let complement = target - num;
        if let Some(&comp_idx) = seen.get(&complement) {
            return vec![comp_idx, idx];
        }
        seen.insert(num, idx);
    }
    return vec![]
}
