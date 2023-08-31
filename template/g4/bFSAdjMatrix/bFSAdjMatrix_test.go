package bfsadjmatrix

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	matrix2 := [][]int{
		{0, 3, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}

	expectedResult1 := []int{0, 1, 4, 5, 6}
	result1 := BFS(matrix2, 0, 6)
	if !reflect.DeepEqual(result1, expectedResult1) {
		t.Errorf("Expected %v but got %v", expectedResult1, result1)
	}

	expectedResult2 := []int(nil)
	result2 := BFS(matrix2, 6, 0)
	if !reflect.DeepEqual(result2, expectedResult2) {
		t.Errorf("Expected %v but got %v", expectedResult2, result2)
	}
}
