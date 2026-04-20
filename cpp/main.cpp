#include <iostream>
#include <vector>

using std::vector;

vector<int> twoSum(vector<int> nums, int target) {
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

  vector<int> response = twoSum(nums, target);

  std::cout << "Response: " << response << std::endl;
}
