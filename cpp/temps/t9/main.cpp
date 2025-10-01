#include <iostream>
#include <vector>

std::vector<int> getResponse(int target, const std::vector<int>& nums) {

    for (int i = 0; i < nums.size(); ++i) {
        for (int j = i + 1; j < nums.size(); ++j)
            if (nums[i] + nums[j] == target) {
                return {i, j};
            }
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
