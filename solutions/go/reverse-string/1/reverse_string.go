package reverse

func Reverse(input string) (result string) {
	for _, r := range input {
        result = string(r) + result
    }
    return result
}
