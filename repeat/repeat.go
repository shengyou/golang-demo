package repeat

func Repeat(count int, char string) string {

	if char == "" {
		char = "z"
	}

	var repeatedString string

	for i := 0; i < count; i++ {
		repeatedString += char
	}

	return repeatedString
}
