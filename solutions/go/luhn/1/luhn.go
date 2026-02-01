package luhn

func Valid(id string) bool {
	var sum int
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
	}

	return sum%10 == 0 && sum != 0
}
