use std::collections::HashMap;

fn main() {
    let nums = vec![1, 2, 3, 4];
    let target = 7;
    let resp = two_sum(nums, target);

    println!("Response: {:?}", resp)
}

pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut seen = HashMap::new();

    for (idx, num) in nums.iter().enumerate() {
        let complement = target - num;
        if let Some(&comp_idx) = seen.get(&complement) {
            return vec![comp_idx as i32, idx as i32];
        }
        seen.insert(num, idx);
    }
    vec![]
}
