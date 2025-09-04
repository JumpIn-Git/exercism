package raindrops

import "strconv"

func Convert(number int) (result string) {
	s := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	for i, p := range s {
		if number%i == 0 {
			result += p
		}
	}
	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}
