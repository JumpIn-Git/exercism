package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}

	var result strings.Builder
	char := rune(input[0])
	count := 1

	for _, r := range input[1:] {
		if r == char {
			count++
		} else {
			if count > 1 {
				result.WriteString(strconv.Itoa(count))
			}
			result.WriteRune(char)
			char = r
			count = 1
		}
	}

	if count > 1 {
		result.WriteString(strconv.Itoa(count))
	}
	result.WriteRune(char)

	return result.String()
}
func RunLengthDecode(input string) string {
	count := ""
	result := ""
	for _, r := range input {
		if r >= '0' && r <= '9' {
			count += string(r)
			continue
		}
		num := 1
		if count != "" {
			num, _ = strconv.Atoi(count)
		}
		result += strings.Repeat(string(r), num)
		count = ""
	}
	return result
}
