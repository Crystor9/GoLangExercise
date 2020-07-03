package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	stringMap := make(map[string]int)
	for i := range arr {
		stringMap[arr[i]] += 1
	}
	return stringMap
}

func main() {
	wc.Test(WordCount)
}
