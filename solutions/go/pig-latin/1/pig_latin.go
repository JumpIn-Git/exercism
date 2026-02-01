package piglatin

import (
	"strings"
)

func Sentence(sentence string) string {
	words := strings.Fields(sentence)

	for i, w := range words {
		// Rule 1
		if vowel(w[0]) || strings.HasPrefix(w, "xr") || strings.HasPrefix(w, "yt") {
			words[i] += "ay"
			continue
		}

		for idx := 0; idx < len(w); idx++ {
			if w[idx] == 'q' && idx+1 < len(w) && w[idx+1] == 'u' {
				words[i] = w[idx+2:] + w[:idx+2] + "ay"
				break
			}

			if (w[idx] == 'y' && idx > 0) || vowel(w[idx]) {
				words[i] = w[idx:] + w[:idx] + "ay"
				break
			}
		}
	}

	return strings.Join(words, " ")
}

func vowel(w byte) bool {
	switch w {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
