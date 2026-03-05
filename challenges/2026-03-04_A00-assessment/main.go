package main

import (
	"fmt"
	"sort"
)

// TopN(nums []int, n int) []int` that returns the top N largest numbers from a slice, **without sorting the original slice**, in descending order.

func main() {

	fmt.Print(TopN([]int{1, 2, 3, 4, 5, 6}, 2))

}

func TopN(nums []int, n int) []int {

	if len(nums) == 0 || n < 0 {
		return []int{}
	}
	if n > len(nums) {
		return nums
	}

	out := make([]int, 0, n)

	for _, val := range nums {

		out = append(out, val)
		sort.Slice(out,
			func(i, j int) bool {
				return out[i] < out[j]
			})
		if len(out) > n {
			out = out[:n]
		}
	}
	return out
}
