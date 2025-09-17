package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() (result uint64) {
	for i := 1; i <= 64; i++ {
		n, _ := Square(i)
		result += n
	}
	return result
}
