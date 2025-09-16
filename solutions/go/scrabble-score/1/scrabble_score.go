package scrabble

import "strings"

func Score(word string) (points int) {
	for _, r := range strings.ToUpper(word) {
        switch r {
        case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
            points += 1
        case 'D', 'G':
            points += 2
        case 'B', 'C', 'M', 'P':
            points += 3
        case 'F', 'H', 'V', 'W', 'Y':
            points += 4
        case 'K':
            points += 5
        case 'J', 'X':
            points += 8
        case 'Q', 'Z':
            points += 10
    	}
    }
    return points
}
