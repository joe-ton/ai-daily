use std::collections::HashMap;

#[derive(Debug, PartialEq)]
enum TwoSumError {
    NoSolution
}

fn main() {
    let nums = vec![1, 2, 3, 4];
    let target = 7;

    match run(&nums, target) {
        Ok((i, j)) => println!("Result: {i}, {j}"),
        Err(e) => eprintln!("Error: {:?}", e),
    }
}

fn run(nums: &[i32], target: i32) -> Result<(usize, usize), TwoSumError> {
    let mut seen = HashMap::new();

    for (idx, num) in nums.iter().enumerate() {
        let complement = target - num;

        if let Some(&comp_idx) = seen.get(&complement) {
            return Ok((comp_idx, idx));
        }

        seen.insert(num, idx);
    }

    Err(TwoSumError::NoSolution)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        let nums = vec![1, 2, 3, 4];
        let target = 7;

        match run(&nums, target) {
            Ok((i, j)) => assert_eq!((i, j), (2, 3)),
            Err(e) => eprintln!("Error: {:?}", e),
        }
    }
}
