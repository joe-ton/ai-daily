
function main(): void {
    const nums = [1, 2, 3, 4];
    const target = 7;
    const result = run(nums, target);

    console.log("Result: ", result);
}

function run(nums: number[], target: number): [number, number] | null {
    if (nums.length < 2) {
        return null
    }

    const seen = new Map<number, number>(); // complements; int to idx

    for (let i = 0; i < nums.length; i++) {
        const num = nums[i];
        const complement = target - num;

        if (seen.has(complement)) {
            const compIdx = seen.get(complement);
            return [compIdx, i];
        }
        seen.set(num, i);
    }

    return null
}

main()
