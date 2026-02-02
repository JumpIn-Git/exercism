package isogram

import "strings"

func IsIsogram(word string) bool {
	m := make(map[rune]struct{}, 26)
	for _, r := range strings.ToLower(word) {
		if r == ' ' || r == '-' {
			continue
		}
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}
