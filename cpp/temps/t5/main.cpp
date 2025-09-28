#include <iostream>
#include <vector>
#include <unordered_map>

std::vector<int> getTwoSum(int target, const std::vector<int, int>& nums) {
    std::unordered_map<int, int> previous;

    for (int i = 0; i < nums.size(); i++) {
        int complement = target - nums[i];

        if (previous.find(complement) != previous.end()) {
            return {previous[complement], i};
        }
        previous[nums[i]] = i;
    }
    return {};
}

int main() {
    int target = 7;
    std::vector<int, int> nums = {1, 2, 3, 4};

    auto result = getTwoSum(target, nums);

    std::cout << "Result: ";

    for (const auto& res : result) {
        std::cout << res << " " << std::endl;
    }
}

