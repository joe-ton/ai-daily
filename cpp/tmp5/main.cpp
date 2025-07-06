#include <iostream>
#include <unordered_map>
#include <vector>

std::pair<int, int> twoSum(const std::vector<int>& nums, int target) {
    if (nums.size() < 2) {
        return {-1, -1};
    }

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
    
    auto result = twoSum(nums, target);

    if (result.first != -1) {
        std::cout << result.first << result.second << std::endl; 
    } else {
        std::cout << "False" << std::endl;
    }

    return 0;
}
