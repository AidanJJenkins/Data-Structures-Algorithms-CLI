package twocrystalballs

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	index := rand.Intn(10000)
	data := make([]bool, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = true
	}

	result := twoCrystalBalls(data)
	if result != index {
		t.Errorf("Expected result: %d, got: %d", index, result)
	}

	// Test with an array of 821 elements filled with false
	result = twoCrystalBalls(make([]bool, 821))
	if result != -1 {
		t.Errorf("Expected result: -1, got: %d", result)
	}
}
