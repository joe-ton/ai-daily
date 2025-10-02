#include <iostream>
#include <vector>
#include <unordered_map>

using std::vector;

vector<int> getResponse(int target, const vector<int>& nums) {
    std::unordered_map<int, int> prev;

    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];
        if (prev.count(complement)) {
            return {prev[complement], i};
        }
        
        prev[nums[i]] = i;
    }

    return {};
}

int main() {
    int target = 7;
    vector<int> nums = {1, 2, 3, 4};

    auto result = getResponse(target, nums);

    std::cout << "Result: ";
    for (const auto& idx : result) {
        std::cout << idx << " ";
    }
}
