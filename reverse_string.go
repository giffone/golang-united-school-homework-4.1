package reverse_string

func ReverseString(input string) (output string) {
	b := []byte{}
	for i := len(input) - 1; i >= 0; i-- {
		b = append(b, input[i])
	}
	return string(b)
}
