package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	input := "hello world"
	output := ReverseWords(input)
	fmt.Println(output)
}

func ReverseWords(input string) string {
	inputSlice := strings.Fields(input)
	i := 0
	j := len(inputSlice) - 1
	for i < j {
		inputSlice[i], inputSlice[j] = inputSlice[j], inputSlice[i]
		i++
		j--
	}
	output := fmt.Sprintf("%s", inputSlice)
	log.Println(output)
	return output
}
