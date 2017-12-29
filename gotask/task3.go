package gotask

func ReverseString(str string) string {
	var temp byte
	bytes := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		bytes[i] = str[i]
	}

	for i := 0; i < len(bytes) / 2; i++ {
		temp = bytes[i]
		bytes[i] = bytes[len(bytes) - i - 1]
		bytes[len(bytes) - i - 1] = temp
	}
	
	return string(bytes)
}