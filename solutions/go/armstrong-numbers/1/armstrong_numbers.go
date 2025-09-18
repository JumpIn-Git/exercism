package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	sum := 0
	for _, r := range s {
		sum += int(math.Pow(float64(r)-'0', float64(len(s))))
	}
	return sum == n
}
