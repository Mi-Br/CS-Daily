package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {

	// // Example: Chunk([]int{1,2,3,4,5}, 2) → [[1 2] [3 4] [5]]
	// // Example: Chunk([]int{1,2,3}, 3)     → [[1 2 3]]
	// // Example: Chunk([]int{}, 2)          → []

	testCases := []struct {
		inp  []int
		size int
		want [][]int
	}{{
		inp:  []int{1, 2, 3, 4, 5},
		size: 2,
		want: [][]int{[]int{1, 2}, []int{3, 4}, []int{5}},
	},
		{
			inp:  []int{1, 2, 3},
			size: 3,
			want: [][]int{[]int{1, 2, 3}},
		},
		{
			inp:  []int{},
			size: 2,
			want: [][]int{},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestCase NR: %d", i+1), func(t *testing.T) {
			got := Chunk(tc.inp, tc.size)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("Incorrect chunks of %d, want: %v, got %v", tc.size, tc.want, got)
			}
		})
	}

}
