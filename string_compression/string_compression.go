package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "aaabbccddd"
	output := CompressString(input)
	fmt.Println(output)
}

func CompressString(input string) string {
	runes := []rune(input)
	fmt.Println(runes[0])
	var out []string
	count := 1
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == runes[i+1] {
			count++
		} else {
			out = append(out, strconv.Itoa(count)+string(runes[i]))
			count = 1
		}
	}
	// Append the count and character of the last group
	out = append(out, strconv.Itoa(count)+string(runes[len(runes)-1]))

	return concat(out)
}

func concat(s []string) string {
	var result string
	for _, str := range s {
		result += str
	}
	return result
}
