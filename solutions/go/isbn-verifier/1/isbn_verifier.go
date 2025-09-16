package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, r := range isbn {
		if i == 9 && r == 'X' {
			sum += 10
			continue
		}
		if r >= '0' && r <= '9' {
			sum += (int(r) - '0') * (10 - i)
		} else {
			return false
		}
	}

	return sum%11 == 0
}
