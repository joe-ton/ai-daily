use std::collections::HashSet;

fn main() {
    let nums = vec![1, 2, 3, 3];
    let result = is_duplicate(&nums);

    println!("Result: {:?}", result);
}

// Determine if given values have duplicate integers
// Return true if so
pub fn is_duplicate(nums: &[i32]) -> bool {
    // hashset
    let mut seen: HashSet<i32> = HashSet::with_capacity(nums.len());

    for &val in nums {
        if seen.contains(&val) {
            return true
        }
        seen.insert(val);
    }
    return false
}

