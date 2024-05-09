package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str1 := "Army"
	str2 := "mary"
	isAnagram := CheckAnagram(str1, str2)
	fmt.Println(isAnagram)

}

func CheckAnagram(str1 string, str2 string) bool {
	//first remove white spaces - convert to lower cases
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	//converts to bytes
	bytes := []byte(str1)
	bytes2 := []byte(str2)

	// Sort the slice of bytes
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	sort.Slice(bytes2, func(i, j int) bool {
		return bytes2[i] < bytes2[j]
	})

	sortedStr1 := string(bytes)
	sortedStr2 := string(bytes2)
	if sortedStr1 == sortedStr2 {
		return true
	}
	return false
}
