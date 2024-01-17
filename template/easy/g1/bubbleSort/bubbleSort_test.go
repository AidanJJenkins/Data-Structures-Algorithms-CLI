package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	foo := []int{9, 3, 7, 4, 72, 700, 42}
	expected := []int{3, 4, 7, 9, 42, 72, 700}

	attempt := bubbleSort(foo)

	if !reflect.DeepEqual(expected, attempt) {
		t.Errorf("Expected: %v, got: %v", expected, attempt)
	}
}
