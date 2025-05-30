function main() {
    let nums = [1, 2, 3, 4];
    let target = 7;
    let resp = run(nums, target);
    console.log("Response: ", resp);
}

function run(nums: number[], target: number): [number, number] {
    const seen = new Map<number, number>();

    for (const [index, value] of nums.entries()) {
        console.log(index, value);
    }
}

main()
