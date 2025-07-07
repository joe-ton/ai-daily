#include <iostream>
#include <vector>
#include <unordered_map>

std::pair<int, int> run(const std::vector<int>& nums, int target) {
    std::unordered_map<int, int> seen;

    for (int i = 0; i < nums.size(); i++) {
        int complement = target - nums[i];
        if (seen.find(complement) != seen.end()) {

        }
    }
}
