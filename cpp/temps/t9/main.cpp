#include <iostream>
#include <vector>
#include <unordered_map>

// Return: two indices
std::pair<int, int> getResponse(int target, const std::vector<int>& nums) {
    // num to idx
    std::unordered_map<int, int> previous_values;

    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];

        if (previous_values.find(complement) != previous_values.end()) {
            return {previous_values[complement], i};
        }

        previous_values[nums[i]] = i;
    }
    return {0, 0};
}

int main() {
    int target = 7;
    std::vector<int> nums = {1, 2, 3, 4};

    auto result = getResponse(target, nums);

    std::cout << "Result: ";

    for (const auto& r : result) {
        std::cout << r << " ";
    }
}
