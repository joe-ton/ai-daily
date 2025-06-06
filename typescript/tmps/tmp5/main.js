function main() {
    var nums = [1, 2, 3, 4];
    var target = 7;
    var resp = run(nums, target);
    console.log("Response: ", resp);
}
function run(nums, target) {
    var seen = new Map();
    for (var i = 0; i < nums.length; i++) {
        var num = nums[i];
        var complement = target - num;
        if (seen.has(complement)) {
            var compIdx = seen.get(complement);
            return [compIdx, i];
        }
        seen.set(num, i);
    }
    return null;
}
main();
