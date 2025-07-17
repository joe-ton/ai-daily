#include <iostream>
#include <unordered_map>
#include <vector>

std::pair<int, int> run(const std::vector<int>& nums, int target) {
    std::unordered_map<int, int> seen;

    for (int i = 0; i < nums.size(); i++) {
        int num = nums[i];
        int complement = target - num;

        if (seen.find(complement) != seen.end()) {
            int compIdx = seen[complement];
            return {compIdx, i};
        }

        seen[num] = i;
    }

    return {-1, -1};
}

int main() {
    std::vector<int> nums = {1, 2, 3, 4};
    int target = 7;

    auto result = run(nums, target);

    if (result.first != -1) {
        std::printf("Result: ", result);
    } else {
        std::printf("False");
    }
}
