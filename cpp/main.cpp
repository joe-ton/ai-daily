#include <iostream>
#include <vector>

using std::vector;

std::vector<int> twoPointers(std::vector<int> nums, int target) {
  int left = 0, right = nums.size() - 1;

  while (left < right) {
    long long sum = (long long)nums[left] + nums[right];

    if (sum == target) {
      return {left + 1, right + 1};
    } else if (sum < target) {
      left++;
    } else {
      right--;
    }
  }
  return {};
}

int main() {
  vector<int> nums = {1, 2, 3, 4};
  int target = 7;

  std::vector<int> response = twoPointers(nums, target);

  std::cout << "Response: ";
  for (const auto &idx : response) {
    std::cout << idx << " ";
  }
  std::cout << std::endl;
}
