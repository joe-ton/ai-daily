function main() {
    var nums = [1, 2, 3, 4];
    var target = 7;
    var resp = run(nums, target);
    console.log(resp);
}
function (nums, target) {
    var seen = new Map();
    for (var i = 0; i < nums.length; i++) {
        var complement = target - nums[i];
        if (seen.has(complement)) {
            compIdx = seen[complement];
            return [compIdx, idx];
        }
        seen[nums[i]] = idx;
    }
    return [];
}
main();
