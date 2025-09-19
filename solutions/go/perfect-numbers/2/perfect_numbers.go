package perfect

import "errors"

// Define the Classification type here.
type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	sum := int64(0)
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	if sum == n {
		return ClassificationPerfect, nil
	} else if sum > n {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}
}
