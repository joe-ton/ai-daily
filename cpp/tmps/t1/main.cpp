#include <iostream>
#include <vector>
#include <unordered_map>

std::vector<int> getResponse(int target, const nums) {
    // num to idx
    std::unordered_map<int, int> prev;

    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];
    }
}

int main() {
    int target = 7;
    std::vector<int> nums = {1, 2, 3, 4};

    std::vector<int> result = getResponse(target, nums);

    std::cout << "Result: ";

    for (const auto& item : result) {
        std::cout << item << " ";
    }
}
