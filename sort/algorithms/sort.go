package sort

import "log"

/**
* TODO: Implement basic sorts
* - Bubble
* - Insertion
* - Selection
* - Quicksort
* - Heapsort
* - Counting sort
* - Radix sort
**/

/**
* TODO: Implement basic searches
* - Binary
* - A*
**/

// TODO: Implement CLI to interact with the algorithms
// TODO: provide a fake data set or generate my own and pipe it into the program
// TODO: sorted items should be green, unsorted red, continually
// TODO: figure out how to output to the same terminal line
// 	- https://github.com/buger/goterm
// 	- https://github.com/ttacon/chalk
// TODO: figure out how to output in different colours

// Bubble will execute a bubble sort
// takes an array of unsorted integers and sorts them
// returns the number of operations used
// sorts in ascending order, smallest on the left, largest right
func Bubble(data []int) (int, int, int, error) {
	log.Println("Bubble Sort")
	log.Println("==========================")
	log.Printf("Unsorted data %v", data)
	temp := data[:]
	iterations, comparisons := 0, 0
	end := (len(data) - 1)
	sorted := false
	for !sorted {
		sorted = true
		iterations++
		for index, curr := range data {
			next := index + 1
			comparisons++
			if index < end && curr > temp[next] {
				temp[index], temp[next] = temp[next], curr
				if sorted {
					sorted = false
				}
			}
		}
	}
	log.Printf("Sorted data %v", temp)
	return iterations, comparisons, len(data), nil
}

// Execute will run the sort
func Execute() {
	testData := []int{0, 5, 3, 7, 4, 12, 1, 8, 11, 2, 6, 9, 10}
	iterations, comparisons, items, err := Bubble(testData)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("%d operations through %d items with %d comparisons made", iterations, items, comparisons)
	}
}
