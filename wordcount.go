package main

import (
	"strings"

	"golang.org/x/tour/wc"
	//"fmt"
)

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	var wc = make(map[string]int)
	for _, str := range strs {
		if _, ok := wc[str]; ok {
			wc[str]++
		} else {
			wc[str] = 1
		}
	}
	return wc
}

func main() {
	//	s := "this is a string is a string"
	//	WordCount(s)
	wc.Test(WordCount)
}
