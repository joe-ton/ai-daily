#include <iostream>
#include <vector> 
#include <unordered_map>

std::pair<int, int> twoSum(const std::vector<int>& nums, int target) {
    std::unordered_map<int, int> seen;

    for (int i = 0; i < nums.size(); i++) {
        int complement = target - nums[i];

        if (seen.find(complement) != seen.end()) {
            return {seen[complement], i};
        }

        seen[nums[i]] = i;
    }

    return {-1, -1};
}

int main() {
    std::vector<int> nums = {1, 2, 3, 4};
    int target = 7;

    std::pair<int, int> result = twoSum(nums, target);
    if (result.first != -1) {
        std::cout << "Indices: " 
            << result.first << ", " 
            << result.second << std::endl;
    } else {
        std::cout << "No solution" << std::endl;
    }

    return 0;
}
