package leetcode

import "testing"

func twoSum(nums []int, target int) []int {
	var numMap = make(map[int]int)
	for index1, num := range nums {
		index2, ok := numMap[target-num]
		if ok {
			return []int{index2, index1}
		}

		numMap[num] = index1
	}

	return nil
}

func Test_twoSum(t *testing.T) {
	t.Run("", func(t *testing.T) {
		got := twoSum([]int{2, 7, 11, 15}, 9)
		t.Log(got)
	})
}
