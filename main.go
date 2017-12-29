package main

import (
    "fmt"
    "softgo/gotask"
    "softgo/gotask/battleShip"
)

func main() {

    //task1
    gotask.PrintString("World")

    //task2
    input := []float32{1, 2, 3, 4}
    fmt.Println(gotask.Sum(input))
    fmt.Println(gotask.Multiply(input))

    //task3
    inputString := "Hello, World!"
    fmt.Println(gotask.ReverseString(inputString))

    //task4
    inputPalindromeString := "radar"
    fmt.Println(gotask.IsPalindrome(inputPalindromeString))

    //task5
    inputHistogramData := []int{4, 9, 7}
    gotask.Histogram(inputHistogramData)

    //task6
    inputStringToEncrypt := "abcd"
    fmt.Println(gotask.CaesarCipher(inputStringToEncrypt, 27))

    //task7
    matrix := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
    fmt.Println(gotask.DiagonalReverse(matrix))

    //task8
    gotask.Game(1, 5)

    //task9
    balancedBrackets := "[[][]]"
    unbalancedBrackets := "[]][[]"
    fmt.Println(gotask.IsBalanced(balancedBrackets))
    fmt.Println(gotask.IsBalanced(unbalancedBrackets))

    //task10
    charsToCount := "sddddwwwssaaaasxccxcdeedddsssaaass"
    fmt.Println(gotask.CharFrequency(charsToCount))

    //task11
    fmt.Println(gotask.Decimal2Binary(10))

	//task12
    battleShip.PlayShipButtleground(10)
}