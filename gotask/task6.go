package gotask

func CaesarCipher(str string, key int) string {

	var normKey int

	if key > 26 {
		normKey = key % 26
	} else if key >= 0 && key <= 26 {
		normKey = key
	} else {
		normKey = 1
	}

	bytes := make([]byte, len(str))

	for i := 0; i < len(bytes); i++ {
		bytes[i] = byte(str[i]) + byte(normKey)
	}
	
	return string(bytes)
}