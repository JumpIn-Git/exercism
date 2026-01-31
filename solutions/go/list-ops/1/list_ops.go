package listops

import "slices"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, i := range s {
		acc = fn(acc, i)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := make(IntList, 0, s.Length())
	for _, i := range s {
		if fn(i) {
			res = append(res, i)
		}
	}
	return res
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	l := make(IntList, 0, s.Length())

	for _, i := range s {
		l = append(l, fn(i))
	}

	return l
}

func (s IntList) Reverse() IntList {
	slices.Reverse(s)
	return s
}

func (s IntList) Append(lst IntList) IntList {
	for _, i := range lst {
		s = append(s, i)
	}

	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = append(s, l...)
	}

	return s
}
