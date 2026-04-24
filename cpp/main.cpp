#include <iostream>
#include <unordered_map>
#include <vector>

using std::vector;

vector<int> twoSum(vector<int> nums, int target) {
  std::unordered_map<int, int> prev;

  for (const auto &idx : nums) {
    int complement = target - nums[idx];

    if (prev.find(complement) != prev.end()) {
      return {prev[complement], idx};
    }
    prev[nums[idx]] = idx;
  }
  return {};
}
