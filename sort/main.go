package main

import (
	"log"

	"github.com/theatlasroom/go-learn/sort/algorithms"
)

func main() {
	log.Println("BEGIN")

	testData := []int{0, 5, 3, 7, 4, 12, 1, 8, 11, 2, 6, 9, 10}
	completed, err := sort.Sort(testData, sort.Bubble)
	if err != nil {
		log.Println(err)
	}

	log.Println("END", completed)
}
