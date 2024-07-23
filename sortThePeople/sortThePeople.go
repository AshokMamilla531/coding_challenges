package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	height int
}

type ByHeight []Person

func (a ByHeight) Len() int {
	return len(a)
}

func (a ByHeight) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByHeight) Less(i, j int) bool {
	return a[i].height > a[j].height
}

func sortPeople(names []string, heights []int) []string {
	n := len(names)
	people := make([]Person, n)

	for i := 0; i < n; i++ {
		people[i] = Person{name: names[i], height: heights[i]}
	}
	sort.Sort(ByHeight(people))
	sortedNames := make([]string, n)
	for i, person := range people {
		sortedNames[i] = person.name
	}
	return sortedNames
}
func main() {
	names := []string{"Ashok", "Shilpa", "Kishore"}
	heights := []int{160, 170, 190}
	sortedNames := sortPeople(names, heights)
	fmt.Println("sorted Names:", sortedNames)
}
