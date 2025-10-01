#include <iostream>
#include <vector>

std::vector<int> getResponse(int target, const std::vector<int>& nums) {
    std::unordered_map<int, int>& nums;

    
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
