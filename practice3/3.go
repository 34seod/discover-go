package practice3

import (
	"math"
)

func BinarySearch(src []string, target string) bool {
	var size, index int

	for len(src) != 1 {
		size = len(src) - 1
		index = int(math.Floor(float64(size / 2)))

		if src[index] > target {
			src = src[:index]
		} else if src[index] == target {
			return true
		} else {
			index++
			src = src[index:]
		}
	}
	return src[0] == target
}
