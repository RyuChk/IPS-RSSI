package converter

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func ToIntArray[T Number](input []T) []int {
	buffer := make([]int, len(input))
	for i, v := range input {
		buffer[i] = int(v)
	}

	return buffer
}
