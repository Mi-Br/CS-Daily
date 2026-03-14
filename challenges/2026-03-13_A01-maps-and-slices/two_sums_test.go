package main

import (
	"testing"
)

func TestTwoSums(t *testing.T) {
	testCases := []struct {
		inp     []int
		sum     int
		answers []int
	}{{
		inp:     []int{2, 7, 11, 15},
		sum:     9,
		answers: []int{0, 1},
	}, {
		inp:     []int{3, 2, 4},
		sum:     6,
		answers: []int{1, 2},
	}}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			gotA, gotB := TwoSum(tc.inp, tc.sum)
			wantA, wantB := tc.answers[0], tc.answers[1]
			if gotA != wantA || gotB != wantB {
				t.Errorf("failed sum for %v - got %d,%d, want %d, %d", tc.inp, gotA, gotB, wantA, wantB)
			}
		})
	}
}
