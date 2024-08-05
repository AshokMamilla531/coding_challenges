package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	value int
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{value: 52}
	})
	return instance
}
func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("Both instances are the same")
	} else {
		fmt.Println("Both instances are different")
	}
	fmt.Printf("Value of s1:%d\n", s1.value)
	fmt.Printf("Value of s2:%d\n", s2.value)

}
