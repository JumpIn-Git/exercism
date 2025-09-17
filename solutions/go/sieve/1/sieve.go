package sieve

func Sieve(limit int) (result []int) {
	primes := make([]bool, limit+1)
	primes[0], primes[1] = true, true
	for i := 2; i <= limit; i++ {
		if !primes[i] {
			result = append(result, i)
			for p := i * i; p <= limit; p += i {
				primes[p] = true
			}
		}
	}
	return result
}
