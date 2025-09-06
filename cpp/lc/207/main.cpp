#include <vector>
#include <iostream>

int main() {
    // std::vector<int> numCourses = {0, 1, 2};
    
    std::vector<int> nums = {1, 2, 3};

    for (const auto& num : nums) {
        std::cout << num << " ";
    }
    std::cout << std::endl;
    return 0;
}
