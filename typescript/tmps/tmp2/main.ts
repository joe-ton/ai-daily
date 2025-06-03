function main() {
    const nums = [1, 2, 3, 4];
    const target = 7;
    let resp = run(nums, target)
    console.log(resp)
}

function run(nums: number[], target: number): [number, number] {
    const seen = new Map<number, number>();

    for (let i = 0; i < nums.length; i++) {
        const num = nums[i];
        const complement = target - num

        if (seen.has(complement)) {
            complIdx = seen[complement];
            return [complIdx, i];
        }
        seen.set(num, i)
    }
    return []
}

main();
