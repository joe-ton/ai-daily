#include <iostream>
#include <vector>

// return indices for two values equal to target
std::vector<int> getResponse(int target, const std::vector<int>& nums) {

    // int to idx
    std::unordered_map<int, int> nums;

    // prev + current
    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];

        if 
    }
}

int main() {
    int target = 7; // two variables equals to sum
    std::vector<int> nums = {1, 2, 3, 4};

    std::vector<int> result = getResponse(target, vector);

    std::cout << "Result: ";
    for (const auto& idx : result) {
        std::cout << idx << " ";
    }
}
