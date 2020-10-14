package lc35

import "testing"

func TestSearchInsert(t *testing.T) {
	doTest([]int{1}, 0, 0, t)
	doTest([]int{1}, 1, 0, t)
	doTest([]int{1, 3, 5, 6}, 0, 0, t)
	doTest([]int{1, 3, 5, 6}, 7, 4, t)
	doTest([]int{1, 3, 5, 6}, 2, 1, t)
	doTest([]int{1, 3, 5, 6}, 5, 2, t)
}

func doTest(nums []int, target int, expect int, t *testing.T) {
	if actual := searchInsert(nums, target); actual != expect {
		t.Errorf("nums = %v, target = %v, expect = %v, actual = %v",
			nums, target, expect, actual)
	}
}
