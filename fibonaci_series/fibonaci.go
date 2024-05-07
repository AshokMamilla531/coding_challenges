package main

import "fmt"

func main() {
	max := 10
	Fibonaci(max)
}

func Fibonaci(m int) {
	fmt.Println(0)
	fmt.Println(1)
	a := 0
	b := 1
	for i := 2; i <= m; i++ {
		next := a + b
		fmt.Println(next)
		a, b = b, next
	}
}
