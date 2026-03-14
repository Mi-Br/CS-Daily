package main

func TwoSum(nums []int, target int) (int, int) {

	seen := make(map[int]int)

	for i, v := range nums {
		need := target - v
		index, exists := seen[need]
		if exists {
			return index, i
		}
		seen[v] = i
	}
	return 0, 0
}
