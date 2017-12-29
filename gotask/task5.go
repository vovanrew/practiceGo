package gotask

import "fmt"

func printStars(length int) {
	for i := 0; i < length; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func Histogram(data []int) {
	for _, value := range(data) {
		printStars(value)
	}
}