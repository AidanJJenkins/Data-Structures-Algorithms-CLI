package twocrystalballs

import (
	"math"
)

func twoCrystalBalls(breaks []bool) int {
	distance := len(breaks)
	step := int(math.Floor(math.Sqrt(float64(distance))))

	i := step
	for i < distance {
		if breaks[i] {
			break
		}
		i += step
	}

	i -= step
	for j := 0; j <= step && i < distance; j++ {
		if breaks[i] {
			return i
		}
		i++
	}
	return -1
}
