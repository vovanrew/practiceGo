package gotask

import (
	"math/rand"
	"fmt"
	"time"
)

func GenerateRandWithRange(startRange int, endRange int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(endRange - startRange) + startRange
}

func Game(startRange int, endRange int) {
	random := GenerateRandWithRange(startRange, endRange)

	var input int

	fmt.Printf("Enter your digit in range [%d, %d]: ", startRange, endRange - 1)

	_, err := fmt.Scanf("%d", &input)

	for input != random && err == nil {
		fmt.Println("Try again!")
		_, err = fmt.Scanf("%d", &input)
	}

	fmt.Println("Congratulations!")
}