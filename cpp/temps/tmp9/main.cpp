#include <iostream>
#include <vector>
#include <unordered_map>

std::pair<int, int> run(const std::vector<int>& nums, int target) {
    std::unordered_map<int, int> seen;

    for (int i = 0; i < nums.size(), i++) {
        int complement = target - nums[i];

        if (seen.find(complement) != seen.end()) {
            int compIdx = seen[complement];
            return {compIdx, i};
        }

        seen[nums[i]] = i;
    }

    return {-1, -1};
}

int main() {
    std::vector<int> nums = {1, 2, 3, 4};
    int target = 7;


}
