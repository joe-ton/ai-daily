#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        let nums = vec![1, 2, 3, 4];
        let target = 7;
        let result = run(&nums, target);

        let expected = vec![2, 3];

        assert_eq!(result, expected);
    }
}
