use std::collections::HashMap;

fn run(nums: &[i32], target: i32) -> Vec<usize> {
    if nums.len() < 2 {
        return vec![];
    }
    let mut seen = HashMap::new(); // tracker of previous values, indices
                                   // past:

    // current
    for (idx, num) in nums.iter().enumerate() {
        let complement = target - num;
        if let Some(&comp_idx) = seen.get(&complement) {
            return vec![comp_idx, idx];
        }
        seen.insert(num, idx);
    }

    vec![]
}
