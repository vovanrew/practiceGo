package gotask

func CharFrequency(str string) map[string]int {
	outputMap := make(map[string]int)
	
	for _, r := range(str) {
		outputMap[string(r)]++
	}

	return outputMap
}