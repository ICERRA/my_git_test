package leetcode

func twoSum(nums []int, target int) []int {
	var numMap map[int]int
	for index1, num := range nums {
		index2, ok := numMap[target-num]
		if ok {
			return []int{index2, index1}
		}

		numMap[num] = index1
	}

	return nil
}
