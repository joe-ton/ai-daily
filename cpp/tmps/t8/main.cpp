#include <iostream>
#include <vector>
#include <unordered_map>
#include <chrono>

std::vector<int> getResponse(int target, const std::vector<int>& nums) {
    std::unordered_map<int, int> prev;
    prev.reserve(nums.size());

    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];

        if (prev.count(complement)) {
            return {prev[complement], i};
        }

        prev[nums[i]] = i;
    }

    return {};
}

int main() {
    const int iterations = 50000;
    auto total_duration = std::chrono::nanoseconds::zero();
    for (int i = 0; i < iterations; ++i) {
            auto start = std::chrono::high_resolution_clock::now();
            auto result = getResponse(7, {1, 2, 3, 4});
            auto end = std::chrono::high_resolution_clock::now();
            total_duration += std::chrono::duration_cast<std::chrono::nanoseconds>(end - start);
        }
        auto avg_duration = total_duration / iterations;
        std::cout << "Average duration: " << avg_duration.count() << " ns" << std::endl;
}
