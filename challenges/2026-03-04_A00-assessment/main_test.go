package main

import (
	"testing"
)

func TestTopN(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		n       int
		wantLen int
	}{
		{"empty input", []int{}, 0, 0},
		{"negative input", []int{3, 4}, -2, 0},
		{"normal case ", []int{30, 40, 50}, 2, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TopN(tc.input, tc.n)
			if len(got) != tc.wantLen {
				t.Errorf("got len %d, want %d", got, tc.wantLen)
			}
		})
	}

}

// func TestInvalidInput(t *testing.T) {
// 	inp := []int{}
// 	got := TopN(inp, 0)
// 	if len(got) != 0 {
// 		t.Errorf("TopN(%v, 0) returned slice of  %d , want 0", inp, len(got))
// 	}
// }
