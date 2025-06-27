"use strict";
function main() {
    const nums = [1, 2, 3, 4];
    const target = 7;
    const result = run(nums, target);
    console.log("Result: ", result);
}
function run(nums, target) {
    if (nums.length < 2) {
        return [];
    }
    const seen = new Map();
    for (let i = 0; i < nums.length; i++) {
        const num = nums[i];
        const complement = target - num;
        if (seen.has(complement)) {
            const compIdx = seen.get(complement);
            return [compIdx, i];
        }
        seen.set(num, i);
    }
    return [];
}
main();
