package prime

import (
    "math"
    "errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
        return 0, errors.New("")
    }

    var result int
    var count int
    for i := 2; count < n; i++ {
        if is_prime(i) {
            result = i
            count++
        }
    }

    return result, nil
}

func is_prime(n int) bool {
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}