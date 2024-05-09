package main

import "fmt"

func main() {
	input := "madxam"
	isPal := IsPalindrome(input)
	fmt.Println(isPal)
}

func IsPalindrome(input string) bool {
	inputBytes := []byte(input)
	i := 0
	j := len(inputBytes) - 1
	for i, j = 0, len(input)-1; i < len(input)-1 && j > 0; i, j = i+1, j-1 {
		if inputBytes[i] != inputBytes[j] {
			return false
		}
	}
	return true

}
