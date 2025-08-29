#include <iostream>
#include <vector>
#include <unordered_map>

std::pair<int, int> run(int target, const std::vector<int>& nums) {

}

int main() {
    int target = 4;
    std::vector<int> nums = {1, 2, 3, 4};

    std::pair<int, int> result = run(target, nums);
    
    if (result.first != -1) {
        std::cout << "Result: " << result.first << result.second << "\n";
    } else {
        std::cout << "Error: " << result.first << result.second << "\n";
    }
}
