#include <iostream>
#include <vector> 
#include <unordered_map>

using namespace std;

pair<int, int> twoSum(const vector<int>& nums, int target) {
    unordered_map<int, int> seen;

    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];
        if (seen.count(complement)) {


    }
}

