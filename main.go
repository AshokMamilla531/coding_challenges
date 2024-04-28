package main

import (
	"fmt"
	"github.com/coding_challenges/reverse_array"
)

func main() {
	fmt.Println("each challenge has written as package and contains code coverage")
	input := []int{3, 4, 6, 8, 3, 67, 0}
	revArry := reverse_array.ReverseArray(input)
	fmt.Println(revArry)
}
