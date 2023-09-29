package quicksort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	foo := []int{9, 3, 7, 4, 69, 420, 42}
	expected := []int{3, 4, 7, 9, 42, 69, 420}

	quickSort(foo)

	if !reflect.DeepEqual(expected, foo) {
		t.Errorf("Expected: %v, got: %v", expected, foo)
	}
}
