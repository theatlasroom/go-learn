package main

import (
	"log"

	"github.com/theatlasroom/go-learn/sort/algorithms"
)

var (
	testData = []int{0, 5, 3, 7, 4, 12, 1, 8, 11, 2, 6, 9, 10}
)

func main() {
	log.Println("BEGIN")
	var err error
	var data []int

	data = make([]int, len(testData))
	copy(data, testData)

	_, err = sort.Sort(sort.Bubble, data, true)
	if err != nil {
		log.Println(err)
	}

	data = make([]int, len(testData))
	copy(data, testData)
	_, err = sort.Sort(sort.Selection, data, true)
	if err != nil {
		log.Println(err)
	}

	log.Println("END")
}
