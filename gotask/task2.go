package gotask

func Sum(array []float32) float32 {
	var result float32
	for _, value := range(array) {
		result += value
	}
	
	return result
}

func Multiply(array []float32) float32 {
	var result float32 = 1
	for _, value := range(array) {
		result *= value
	}

	return result
}