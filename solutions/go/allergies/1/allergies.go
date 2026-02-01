package allergies

var list = []struct {
	score uint
	name  string
}{
	{128, "cats"},
	{64, "pollen"},
	{32, "chocolate"},
	{16, "tomatoes"},
	{8, "strawberries"},
	{4, "shellfish"},
	{2, "peanuts"},
	{1, "eggs"},
}

func Allergies(allergies uint) []string {
	var res []string
	for _, item := range list {
		if allergies&item.score != 0 {
			res = append(res, item.name)
		}
	}
	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	for _, item := range list {
		if item.name == allergen {
			if allergies&item.score != 0 {
				return true
			}
			return false
		}
	}
	return false
}
