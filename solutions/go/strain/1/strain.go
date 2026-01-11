package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](s []T, f func(T) bool) (result []T) {
    for _, i := range s {
        if f(i) {
            result = append(result, i)
        }
    }
    return result
}

func Discard[T any](s []T, f func(T) bool) []T {
    return Keep(s, func(x T) bool {
        return !f(x)
    })
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
