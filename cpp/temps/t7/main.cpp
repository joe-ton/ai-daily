#include <iostream>
#include <array>
#include <unordered_map>

std::array<int, 2> getResponse(int target, const std::array<int, 4>& nums) {
    std::unordered_map<int, int> prev;

    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];

        if (prev.find(complement) != prev.end()) {
            return {prev[complement], i};
        }

        prev[nums[i]] = i;
    }

    return {-1, -1};
}

int main() {
    int target = 7;
    std::array<int, 4> nums = {1, 2, 3, 4};

    auto result = getResponse(target, nums);

    std::cout << "Result: ";
    for (const auto& idx : result) {
        std::cout << idx << " ";
    }

}
