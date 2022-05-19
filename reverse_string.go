package reverse_string

func ReverseString(input string) (output string) {
	r := []rune(input)
	lR := len(r)
	s := make([]rune, lR)
	i := 0
	for index := lR - 1; index >= 0; index-- {
		s[i] = r[index]
		i++
	}
	return string(s)
}

// func ReverseString(input string) (output string) {
// 	b := []byte{}
// 	for i := len(input) - 1; i >= 0; i-- {
// 		b = append(b, input[i])
// 	}
// 	return string(b)
// }

// func ReverseString(input string) (output string) {
// 	b := make([]byte, len(input))
// 	i := 0
// 	for index := len(input) - 1; index >= 0; index-- {
// 		b[i] = input[index]
// 		i++
// 	}
// 	return string(b)
// }
