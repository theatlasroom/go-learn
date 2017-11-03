package main

import (
	"log"

	"github.com/theatlasroom/go-learn/sort/algorithms"
)

func main() {
	log.Println("BEGIN")

	testData := []int{0, 5, 3, 7, 4, 12, 1, 8, 11, 2, 6, 9, 10}
	sort.Sort(testData, sort.Bubble)

	log.Println("END")
}
