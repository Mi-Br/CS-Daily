package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	testCases := []struct {
		inp  []string
		want map[string]int
	}{{
		inp:  []string{"go", "is", "go", "Go", "fast"},
		want: map[string]int{"go": 2, "Go": 1, "is": 1, "fast": 1},
	}}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {

			got := WordFrequency(tc.inp)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Word frequency does not match: got %v, want %v", got, tc.want)
			}
		})
	}

}
