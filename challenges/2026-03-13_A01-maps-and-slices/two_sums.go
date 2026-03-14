package main

func TwoSum(nums []int, target int) (int, int) {
	// Brute forse solution, its O(n*n)
	// for i := 0; i < len(nums)-1; i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			return i, j
	// 		}
	// 	}
	// }
	// return 0, 0

	seen := make(map[int]int)

	for i, v := range nums {
		need := target - v
		index, exists := seen[need]
		if exists {
			return ascOrder(i, index)
		} else {
			seen[v] = i
		}
	}
	return 0, 0
}

func ascOrder(x, y int) (int, int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}
