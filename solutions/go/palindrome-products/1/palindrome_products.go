package palindrome

import (
	"errors"
)

// Define Product type here.
type Product struct {
	Value          int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	var smallest, biggest *Product
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			prod := i * j
			if !isPalindrome(prod) {
				continue
			}

			if smallest == nil || prod < smallest.Value {
				smallest = &Product{Value: prod, Factorizations: [][2]int{{i, j}}}
			} else if prod == smallest.Value {
				smallest.Factorizations = append(smallest.Factorizations, [2]int{i, j})
			}

			if biggest == nil || prod > biggest.Value {
				biggest = &Product{Value: prod, Factorizations: [][2]int{{i, j}}}
			} else if prod == biggest.Value {
				biggest.Factorizations = append(biggest.Factorizations, [2]int{i, j})
			}
		}
	}

	if smallest == nil || biggest == nil {
		return Product{}, Product{}, errors.New("no palindromes")
	}

	return *smallest, *biggest, nil
}

func isPalindrome(n int) bool {
	orig := n
	var reversed int
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}

	return orig == reversed
}
