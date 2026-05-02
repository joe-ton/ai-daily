#include <iostream>
#include <vector>

using std::vector;

vector<int> twoSum(vector<int> nums, int target) {
  int left_idx = 0, right_idx = nums.size() - 1;

  while (left_idx < right_idx) {
    int sum = nums[left_idx] + nums[right_idx];

    if (sum == target) {
      return {left_idx, right_idx};
    } else if (sum < target) {
      left_idx++;
    } else {
      right_idx--;
    }
  }
  return {};
}

int main() {
  vector<int> nums = {1, 2, 3, 4};
  int target = 7;

  vector<int> response = twoSum(nums, target);

  std::cout << "Response: ";
  for (const auto &idx : response) {
    std::cout << idx << " ";
  }
  std::cout << std::endl;
}
