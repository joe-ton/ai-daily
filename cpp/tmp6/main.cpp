#include <iostream>
#include <vector>
#include <unordered_map>

std::pair<int, int> run(const std::vector<int>& nums, int target) {
    if (nums.size() < 2) {
        return {-1, -1};
    }

    std::unordered_map<int, int> seen;

    for (int idx = 0; idx < nums.size(); idx++) {
        int num = nums[idx];
        int complement = target - num;

        if (seen.find(complement) != seen.end()) {
            int compIdx = seen[complement]; 
            return {compIdx, idx};
        }

        seen[num] = idx;
    }

    return {-1, -1};
}

int main() {
    std::vector<int> nums = {1, 2, 3, 4};
    int target = 7;

    std::pair<int, int> result = run(nums, target);

    if (result.first != -1) {
        std::cout << result.first << result.second << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }

    return 0;
}
