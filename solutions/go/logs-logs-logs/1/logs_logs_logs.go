package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
        switch {
            case r == '❗':
            return "recommendation"
            case r == '🔍':
            return "search"
            case r == '☀':
            return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) (result string) {
	for _, r := range log {
        if r == oldRune {
            result += string(newRune)
        } else {
            result += string(r)
        }
    }
    return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    count := 0
    for range log {
        count++
    }
    return count <= limit
}
