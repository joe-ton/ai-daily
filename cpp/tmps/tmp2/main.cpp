#include <iostream>
<<<<<<< HEAD
#include <unordered_map>
#include <vector>
#include <cstdio>

std::pair<int, int> run(const std::vector<int>& nums, int target) {
    std::unordered_map<int, int> seen;

    for (int i = 0; i < nums.size(); i++) {
        int num = nums[i];
        int complement = target - num;

        if (seen.find(complement) != seen.end()) {
            int compIdx = seen[complement];
            return {compIdx, i};
        }

        seen[num] = i;
    }

    return {-1, -1};
}

int main() {
    std::vector<int> nums = {1, 2, 3, 4};
    int target = 7;

    std::pair<int, int> result = run(nums, target);

    if (result.first != -1) {
        std::printf("Result: %d %d\n", result.first, result.second);
    } else {
        std::printf("Error: ", -1);
    }
=======

int main() {
    std::string name = "Joe";

    for (int i = 0; i < name.size(); i++) {
        std::cout << name[i] << std::endl;
    }

    for (char c : name) {
        std::printf("Hello, World\n");
    }

    return 0;
>>>>>>> 15868fd53d71407a5cab07df7717291f671ee1df
}
