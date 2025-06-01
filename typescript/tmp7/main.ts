
function main() {
    const nums = [1, 2, 3, 4];
    const target = 7;
    let resp = run(nums, target);
    console.log(resp)
}

function(nums: number[], target: number): [number, number] {
    const seen = new Map<number, number>();
    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];
        if (seen.has(complement)) {
            compIdx = seen[complement]
            return [compIdx, idx]
        }
        seen.set(nums[i], i);
    }
    return [];
}

main();
