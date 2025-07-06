#include <iostream>
#include <vector>
#include <unordered_map>
#include <cassert>

std::pair<int, int> run(const std::vector<int>& nums, int target) {
    std::unordered_map<int, int> seen;

    for (int i = 0; i < nums.size(); i++) {
        int complement = target - nums[i];

        if (seen.find(complement) != seen.end()) {
            int compIdx = seen[complement];
            return {compIdx, i};
        }

        seen[nums[i]] = i;
    }

    return {-1, -1};
}

void test_run() {
    {
        std::vector<int> nums = {1, 2, 3, 4};
        int target = 7;
        std::pair<int, int> result = run(nums, target);

        std::pair<int, int> expected = std::pair<int, int>{2, 3};

        assert((result == expected));
    }
}

int main() {
    test_run();
    // std::vector<int> nums = {1, 2, 3, 4};
    // int target = 7;
    // std::pair<int, int> result = run(nums, target);
    //
    // if (result.first != -1) {
    //     std::cout << result.first << result.second << std::endl;
    // } else {
    //     std::cout << "False" << std::endl;
    // }
}

