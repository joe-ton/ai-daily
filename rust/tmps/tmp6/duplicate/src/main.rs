use std::collections::HashSet;

fn main() {
    let nums = vec![1, 2, 3, 3];
    let result = is_duplicate(&nums);

    println!("Result: {:?}", result);
}

pub fn is_duplicate(nums: &[i32]) -> bool {
    let mut seen: HashSet<i32> = HashSet::with_capacity(nums.len());

    // iterate
    for &val in nums {
        if seen.contains(&val) {
            return true
        }
        seen.insert(val);
    }
    false
}
