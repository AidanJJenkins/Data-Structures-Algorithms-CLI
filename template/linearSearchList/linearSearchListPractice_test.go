package linearsearchlist

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	var tests = []struct {
		haystack []int
		needle   int
		answer   bool
	}{
		{foo, 69, true},
		{foo, 1336, false},
		{foo, 69420, true},
		{foo, 69421, false},
		{foo, 1, true},
		{foo, 0, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.haystack, tt.needle)
		t.Run(testname, func(t *testing.T) {
			ans := linearSearch(tt.haystack, tt.needle)
			if ans != tt.answer {
				t.Errorf("got %t, want %t", ans, tt.answer)
			}
		})
	}
}
