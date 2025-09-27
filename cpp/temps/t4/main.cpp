#include <iostream>
#include <vector>
#include <unordered_map>

std::vector<int, int> getTwoSum(int target, const std::vector<int>& nums) {

    std::vector<int, int> previous;

    for (int i = 0; i < nums.size(); i++) {
        int complement = target - nums[i];

        if ()
    }
}

int main() {
    int target = 7;
    std::vector<int> nums = {1, 2, 3, 4};

    auto result = getTwoSum(target, nums);

    std::cout << "Result: ";
    for (int i = 0; i < result.size(); i++) {
        std::cout << result[i];
        if (i < result.size() - 1) {
            std::cout << ", ";
        }
    }
}
