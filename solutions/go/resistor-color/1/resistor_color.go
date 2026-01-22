package resistorcolor

// Colors returns the list of all colors.
var Size = 10
func Colors() []string {
	return []string{
        "black",
        "brown",
        "red",
        "orange",
        "yellow",
        "green",
        "blue",
        "violet",
        "grey",
        "white",
    }
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for i := 0; i < Size; i++ {
        if Colors()[i] == color {
            return i
        }
    }
    return -1
}
