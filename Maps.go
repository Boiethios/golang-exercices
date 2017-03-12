package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	
	for _, word := range strings.Fields(s) {
		ans[word] = ans[word] + 1
	}
	return ans
}

func main() {
	wc.Test(WordCount)
}

