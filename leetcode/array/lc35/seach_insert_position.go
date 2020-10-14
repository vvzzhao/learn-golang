package lc35

//searchInsert https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	l := 0
	h := len(nums)
	for l < h {
		if nums[l] >= target {
			return l
		} else if nums[h-1] < target {
			return h
		}
		m := (l + h) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			h = m
		} else if nums[m] < target {
			l = m + 1
		}
	}
	return -1
}
