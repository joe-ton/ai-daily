// LC 2824. Count Pairs Whose Sum is Less Than Target
// Time: O(nlogn)
// Space: O(1)

package countpairs

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)

	result := 0
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			result += right - left
			left++
		} else {
			right--
		}
	}
	return result
}
