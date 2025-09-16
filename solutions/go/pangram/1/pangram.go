package pangram

import "strings"

func IsPangram(input string) bool {
	seen := make([]bool, 26)
	for _, r := range strings.ToLower(input) {
		if r >= 'a' && r <= 'z' {
			seen[r-'a'] = true
		}
	}

	for _, b := range seen {
		if !b {
			return false
		}
	}
	return true
}
