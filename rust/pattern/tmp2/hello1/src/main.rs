fn main() {
    println!("Hello, world!");
}

pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut seen = HashMap::new(); // complement: int to idx


    for (idx, num) in nums.iter().enumerate() {
        let complement = target - num;
        if let Some(&compIdx) = seen.get(&complement) {
            return vec![compIdx as i32, idx as i32];
        }
        seen.insert(num, idx);
    }
    vec![]
}

