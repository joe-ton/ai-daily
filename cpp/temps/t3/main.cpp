#include <iostream>
#include <array>
#include <unordered_map>

std::array<int, 2> 

int main() {
    int target = 7;
    int nums[4] = {1, 2, 3, 4};

    auto result = getResponse(target, nums);

    std::cout << "Result: ";
    for (const auto& idx : result) {
        std::cout << idx << " ";
    }
}
