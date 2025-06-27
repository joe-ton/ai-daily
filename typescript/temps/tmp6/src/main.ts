function main(): void {
    const nums: number[] = [1, 2, 3, 4];
    const target: number = 7;
    const result: number[] = run(nums, target)

    console.log("Result: ", result);
}

function run(nums: number[], target: number): number[] {
    if (nums.length < 2) {
        return []
    }

    const seen = new Map<number, number>();

    for (let i = 0; i < nums.length; i++) {
        const num: number = nums[i];
        const complement = target - num;

        if (seen.has(complement)) {
            const compIdx = seen.get(complement)!;
            return [compIdx, i];
        }
        seen.set(num, i);
    }

    return [];
}

main();
