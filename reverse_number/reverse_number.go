package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := -334
	fmt.Println(ReverseNum(input))
	inputString := strconv.Itoa(input)
	fmt.Println(ReverseNumString(inputString))
}
func ReverseNumString(input string) string {
	sign := ""
	if input[0] == '-' {
		sign = "-"
		input = input[1:]
	}

	runes := []rune(input)
	reversed := make([]rune, len(runes))

	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		reversed[i] = runes[j]
	}

	return sign + string(reversed)
}

func ReverseNum(input int) int {
	sign, rem, rev := 1, 0, 0

	if input < 0 {
		sign = -1
		input = input * sign
	}
	for input > 0 {
		rem = input % 10
		rev = rev*10 + rem
		input = input / 10
	}
	return rev * sign
}
