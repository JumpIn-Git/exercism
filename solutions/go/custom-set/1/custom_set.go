package stringset

import (
	"fmt"
	"slices"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	s []string
}

func New() Set {
	return Set{}
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]struct{}, len(strSlice))
	list := make([]string, 0, len(strSlice))
	for _, item := range strSlice {
		if _, exists := allKeys[item]; !exists {
			allKeys[item] = struct{}{}
			list = append(list, item)
		}
	}
	return list
}

func NewFromSlice(l []string) Set {
	return Set{removeDuplicateStr(l)}
}

func (s Set) String() string {
	if len(s.s) == 0 {
		return "{}"
	}
	return fmt.Sprintf(`{"%s"}`, strings.Join(s.s, "\", \""))
}

func (s Set) IsEmpty() bool {
	return len(s.s) == 0
}

func (s Set) Has(elem string) bool {
	return slices.Contains(s.s, elem)
}

func (s *Set) Add(elem string) {
	if !s.Has(elem) {
		s.s = append(s.s, elem)
	}
}

func Subset(s1, s2 Set) bool {
	m := make(map[string]struct{}, len(s2.s))
	for _, s := range s2.s {
		m[s] = struct{}{}
	}

	for _, s := range s1.s {
		if _, found := m[s]; !found {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	m := make(map[string]struct{}, len(s2.s))
	for _, s := range s2.s {
		m[s] = struct{}{}
	}

	for _, s := range s1.s {
		if _, found := m[s]; found {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.s) != len(s2.s) {
		return false
	}

	m := make(map[string]struct{}, len(s2.s))
	for _, s := range s2.s {
		m[s] = struct{}{}
	}

	for _, s := range s1.s {
		if _, found := m[s]; !found {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) (res Set) {
	for _, s := range s1.s {
		if s2.Has(s) {
			res.Add(s)
		}
	}

	return res
}

func Difference(s1, s2 Set) (res Set) {
	for _, s := range s1.s {
		if !s2.Has(s) {
			res.Add(s)
		}
	}

	return res
}

func Union(s1, s2 Set) (res Set) {
	for _, s := range s1.s {
		res.Add(s)
	}

	for _, s := range s2.s {
		res.Add(s)
	}

	return res
}
