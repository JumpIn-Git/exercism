package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("")
	}
	aRunes := []rune(a)
	bRunes := []rune(b)
	distance := 0
	for i, r := range aRunes {
		if r == bRunes[i] {
			distance++
		}
	}
	return distance, nil
}
