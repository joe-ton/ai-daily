function main() {
    const nums = [1, 2, 3, 4];
    const target = 7;
    let resp = run(nums, target)
    console.log(resp)
}

function two_sum(nums: number[], target: number): number {
    if (nums.length) < 2 return [];

    const seen = new Map<number, number>();

    for (let i = 0; i < nums.length; i++) {
        const complement = target - seen[i];

        if (seen.has(complement)) {

        }
    }
}

main()
