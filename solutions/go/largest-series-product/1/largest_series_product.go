package lsproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("span must be smaller than string length")
	}
	var max int64
	for i := 0; i < len(digits)-span+1; i++ {
		var sum int64 = 1
		for _, r := range digits[i : i+span] {
			if r < '0' || r > '9' {
				return 0, errors.New("digits input must only contain digits")
			}
			sum *= int64(r) - '0'
		}
		if sum > max {
			max = sum
		}
	}
	return max, nil
}
