package gotask

func IsBalanced(str string) bool {
	var balanceCounter int

	for _, char := range(str) {
		if string(char) == "[" {
			balanceCounter++
		} else if string(char) == "]" {
			if balanceCounter == 0 {
				return false
			} 
			balanceCounter--
		} else {
			continue
		}
	}

	return balanceCounter == 0
}