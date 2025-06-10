
function main(): void {
    const nums = [1, 2, 3, 3];
    const result = run(nums);

    console.log(result)
}

function run(nums: number[]): boolean {
    const seen = new Set<number>();

    for (let i = 0; i < nums.length; i++) {
        const val = nums[i];

        if (seen.has(val)) {
            return true;
        }
        seen.add(val);
    }
    return false;
}


main();
