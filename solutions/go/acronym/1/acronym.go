// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	f := strings.FieldsFunc(s, func(r rune) bool {
        return r == ' ' || r == '-' || r == '_'
    })

    var res strings.Builder
    for _, w := range f {
        res.WriteString(strings.ToUpper(w[:1]))
    }

    return res.String()
}
