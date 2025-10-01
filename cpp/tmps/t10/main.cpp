#include <iostream>
#include <vector>
#include <unordered_map>

std::vector<int> getResponse(int target, const std::vector<int>& nums) {
    std::unordered_map<int, int> map;

    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];

        if (map.count(complement)) {
            return {map[complement], i};
        }

        map[nums[i]] = i;
    }

    return {};
}

int main() {
    int target = 7;
    std::vector<int> nums = {1, 2, 3, 4};

    auto result = getResponse(target, nums);

    std::cout << "Result: ";
    for (const auto& idx : result) {
        std::cout << idx << " ";
    }
}
