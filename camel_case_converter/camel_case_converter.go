package main

import (
	"log"
	"strings"
)

func main() {
	input1 := "hello_world"
	input2 := "helloworld ashok"
	input3 := "ashok_mamilla-hello_world_code"

	result1 := ToCamelCase(input1)
	result2 := ToCamelCase(input2)
	result3 := ToCamelCase(input3)
	log.Println("Result 1:", result1)
	log.Println("Result 2:", result2)
	log.Println("Result 3:", result3)
}

func ToCamelCase(input string) string {
	// Split the input string into words based on underscores (_) or hyphens (-).
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == '_' || r == '-'
	})

	// Capitalize the first letter of each word after the first word.
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	// Join the words together to form the camelCase string.
	camelCaseString := strings.Join(words, "")

	return camelCaseString
}
