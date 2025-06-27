use std::collections::HashMap;

fn main() {
    let nums = vec![1, 2, 3, 4];
    let target = 7;
    let result = run(&nums, target);

    println!("Result: {:?}", result);
}

fn run(nums: &[i32], target: i32) -> Vec<usize> {
    let mut seen = HashMap::new();

    for (idx, num) in nums.iter().enumerate() {
        let complement = target - num;

        if let Some(&comp_idx) = seen.get(&complement) {
            return vec![num, idx];
        }
        seen.insert(comp_idx, idx);
    }

    vec![]
}

#[cfg(test)]

mod tests {
    use super::*;

    fn test_two_sum() {
        let nums = vec![1, 2, 3, 4];
        let target = target;
        let result = run(&nums, target);

        let expected = vec![2, 3];

        assert_eq!(result, expected)
    }
}
