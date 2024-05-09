package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	input := "The quick brown fox jumps over the lazy dog."
	mapper := WordCount(input)
	for i, j := range mapper {
		fmt.Println(i+":", j)
	}
}
func WordCount(input string) map[string]int {
	lowerInput := strings.ToLower(input)
	log.Println(lowerInput)
	inputSlice := strings.Fields(lowerInput)
	wordCountMap := make(map[string]int)
	for i := 0; i < len(inputSlice); i++ {
		wordCountMap[inputSlice[i]]++
	}
	return wordCountMap
}
