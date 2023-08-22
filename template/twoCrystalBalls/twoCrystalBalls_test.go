package twocrystalballs

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	idx := rand.Intn(10000)
	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	if result := twoCrystalBalls(data); result != idx {
		t.Errorf("Expected %d, but got %d", idx, result)
	}

	if result := twoCrystalBalls(make([]bool, 821)); result != -1 {
		t.Errorf("Expected -1, but got %d", result)
	}
}
