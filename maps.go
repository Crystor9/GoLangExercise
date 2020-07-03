package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	stringMap := make(map[string]int)
	for i := range arr {
		stringMap[arr[i]] += 1
	}
	return stringMap
}

// func main() {
// 	wc.Test(WordCount)
// }
