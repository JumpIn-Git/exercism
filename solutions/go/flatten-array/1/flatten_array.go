package flatten

func Flatten(nested any) []any {
	s := []any{}
	switch i := nested.(type) {
	case []any:
		for _, v := range i {
			s = append(s, Flatten(v)...)
		}
	case any:
		s = append(s, i)
	}
	return s
}
