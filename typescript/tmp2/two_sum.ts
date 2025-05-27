
function run(nums: number[], target: number): number[] {
    const seen = new Map<number, number>();

    for (let i = 0; i < nums.length; i++) {
        const complement =  target - nums[i];
        if (seen.has(complement)) {
            return [seen.get(complement)!, i];
        }
        seen.set(nums[i], i);
    }
    return [];
}

function main() {
    const nums = [1, 2, 3, 4];
    const target = 7;
    const result = run(nums, target);
    console.log(result);
}

main();
