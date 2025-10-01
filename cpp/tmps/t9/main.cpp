#include <vector>
#include <iostream>

int main() {
    int target = 7;
    std::vector<int> nums = {1, 2, 3, 4};

    std::cout << "Vector: ";

    for (const auto& value : nums) {
        std::cout << value << " ";
    }


    return 0;
}
