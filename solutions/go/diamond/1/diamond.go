package diamond

import "strings"

func Gen(char byte) (string, error) {
	if char == 'A' {
		return "A", nil
	}

	spaces := int(char - 'A')
	lines := make([]string, spaces)
	for i := 0; i <= spaces; i++ {
		var line strings.Builder
		line.WriteString(strings.Repeat(string(char), spaces-i))
		line.WriteByte(char)

		if i > 0 {
			line.WriteString(strings.Repeat(" ", (spaces*2+2)-((spaces-i)*2+2)))
			line.WriteByte(char)
		}

		line.WriteString(strings.Repeat(string(char), spaces-i))
		lines = append(lines, line.String())
	}

	var result strings.Builder
	for _, i := range lines {
		result.WriteString(i)
	}
	for i := spaces - 1; i >= 0; spaces-- {
		result.WriteString(lines[i])
	}
	return result.String(), nil
}
