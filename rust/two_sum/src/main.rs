use std::collections::HashMap;

/// Returns indices of the two numbers that add up to the target.
fn two_sum(nums: Vec<i32>, target: i32) -> Option<(usize, usize)> {
    let mut map = HashMap::new(); // value -> index

    for (i, num) in nums.iter().enumerate() {
        let complement = target - num;
        if let Some(&j) = map.get(&complement) {
            return Some((j, i)); // return the pair of indices
        }
        map.insert(num, i);
    }

    None // if no solution found
}

fn main() {
    let nums = vec![2, 7, 11, 15];
    let target = 9;
    match two_sum(nums, target) {
        Some((i, j)) => println!("Indices: {}, {}", i, j),
        None => println!("No solution found."),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum_found() {
        let result = two_sum(vec![2, 7, 11, 15], 9);
        assert_eq!(result, Some((0, 1)));
    }

    #[test]
    fn test_two_sum_not_found() {
        let result = two_sum(vec![1, 2, 3], 7);
        assert_eq!(result, None);
    }

    #[test]
    fn test_two_sum_with_negatives() {
        let result = two_sum(vec![-3, 4, 3, 90], 0);
        assert_eq!(result, Some((0, 2)));
    }
}

