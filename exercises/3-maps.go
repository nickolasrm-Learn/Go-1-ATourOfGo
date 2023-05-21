package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		val, ok := counter[word]
		if ok {
			counter[word] = val + 1
		} else {
			counter[word] = 1
		}
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
