package luhn

func Valid(id string) bool {
	var sum, size int
	double := false
	for i := len(id) - 1; i >= 0; i-- {
		r := id[i]

		if r == ' ' {
			continue
		}

		if r < '0' || r > '9' {
			return false
		}

		n := int(r - '0')
		if double {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		double = !double

		sum += n
		size++
	}

	return sum%10 == 0 && size > 1
}
