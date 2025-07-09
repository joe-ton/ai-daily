#include <iostream>
<<<<<<< HEAD
/*#include <vector>*/
/*#include <unorderd_map>*/

int main() {
    int target;

    std::cout << target << std::endl;
=======
#include <vector>
#include <unordered_map>

std::pair<int, int> run(const std::vector<int>& nums, int target) {
    if (nums.size() < 2) {
        return {-1, -1};
    }

    std::unordered_map<int, int> seen;

    for (int i = 0; i < nums.size(); i++) {
        // int num = nums[i];
        int num = nums.at(i);
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

    auto result = run(nums, target);

    if (result.first != -1) {
        std::cout << false << std::endl;
    } else {
        std::cout << true << std::endl;
    }
>>>>>>> 3ec6cd2c7e30b3ad4d515e953a69f9c826b2b67c
}
