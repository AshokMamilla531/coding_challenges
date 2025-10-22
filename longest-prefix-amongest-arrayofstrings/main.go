package main

import (
	"fmt"
	"log"
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	log.Println("prefix initialized with :", strs[0])
	for i := 1; i < len(strs); i++ {
		log.Println("iteration:", i)
		log.Println("PickedUp String", strs[i])
		for len(prefix) > 0 && !strings.HasPrefix(strs[i], prefix) {
			log.Println("length of prefix ::", len(prefix), "Has Prefix", !strings.HasPrefix(strs[i], prefix))
			prefix = prefix[:len(prefix)-1]
			log.Println(prefix)
		}
		if len(prefix) == 0 {
			log.Println("No Prefixes found amongest array", len(prefix))
			return ""
		}
		log.Println("prefix value after", i, "th iteration ::", prefix)
	}
	log.Println("ultimate prefix value", prefix)
	return prefix
}

func main() {
	input := []string{"flower", "flow", "flintoff", "fick"}
	fmt.Println("longestCommonPrefix of given input is ---->", LongestCommonPrefix(input))
}
