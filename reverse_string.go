package reverse_string

import "bytes"

func ReverseString(input string) (output string) {
	b := bytes.Buffer{}
	for i := len(input) - 1; i >= 0; i-- {
		b.WriteByte(input[i])
	}
	return b.String()
}
