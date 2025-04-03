
function getTwoSum(nums: number[], target: number): [number, number] | null {

    const seen: Record<number, number> = {};

    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];

        if (seen.hasOwnProperty(complement)) {
            return [seen[complement], i];
        }

        seen[nums[i]] = i;
    }

    return null; // Like Go's 'nil' return when no match
}

function main() {
    const nums = [1, 2, 3, 4];
    const target = 7;

    const result = getTwoSum(nums, target);

    if (result === null) {
        console.log("Error: No match found.");
    } else {
        console.log(`Response: [${result[0]}, ${result[1]}]`);
    }
}

main();


