package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hello world world hello"
	count := WordCount(text)
	for i, c := range count {
		fmt.Println(c, i)
	}

	//merge two maps
	m1 := map[string]int{"a": 1, "b": 2, "c": 4}
	m2 := map[string]int{"a": 1, "b": 3, "c": 5}
	m3 := make(map[string]int)
	for i, letter := range m1 {
		m3[i] += letter
	}
	for l, i := range m2 {
		m3[l] += i
	}
	fmt.Println(m3)

	ip := []int{1, 2, 3, 1, 2, 2, 1, 2, 4, 6, 7, 1}
	finder := make(map[int]int)
	for i := 0; i < len(ip); i++ {
		finder[ip[i]]++
	}
	fmt.Println(finder)

	fruits := []string{"apple", "banana", "pear", "grape", "kiwi", "orange"}
	fruitLength := make(map[int][]string)
	for _, fruit := range fruits {
		length := len(fruit)
		fruitLength[length] = append(fruitLength[length], fruit)
	}
	fmt.Println(fruitLength)
}

func WordCount(text string) map[string]int {
	words := strings.Fields(text)
	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[word]++
	}
	return wordCounts
}
