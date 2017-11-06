package sort

import (
	"errors"
	"log"
)

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
// protip: iota increments at each const declaration
const (
	Bubble    = iota
	Selection = iota
	Insertion = iota
)

// Result combines stats and sorted data to provide more information
type Result struct {
	Data                            []int
	Comparisons, Iterations, Length int
	Error                           error
}

func formatResult(
	data []int,
	iterations, comparisons, length int) (Result, error) {
	return Result{
		Data:        data,
		Length:      length,
		Comparisons: comparisons,
		Iterations:  iterations,
		Error:       nil,
	}, nil
}

// func ExtractData

// Bubble sort
// takes an array of unsorted integers and sorts them
// returns the number of operations used
// sorts in ascending order, smallest on the left, largest right
func bubble(data []int, verbose bool) (Result, error) {
	if verbose {
		log.Println("Bubble Sort")
		log.Println("==========================")
		log.Printf("Unsorted data %v", data)
	}
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
	if verbose {
		log.Printf("Sorted data %v", temp)
	}
	return formatResult(temp, iterations, comparisons, len(temp))
}

// Selection sort
// takes an array of unsorted integers and sorts them
// returns the number of operations used
// sorts in ascending order, smallest on the left, largest right
func selection(data []int, verbose bool) (Result, error) {
	if verbose {
		log.Println("Selection Sort")
		log.Println("==========================")
		log.Printf("Unsorted data %v", data)
	}

	// obtain a slice
	// return early if we have 1 or less items
	total := len(data)
	if total <= 1 {
		return formatResult(data, 0, 0, total)
	}

	comparisons, iterations := 0, 0
	sorted := false
	arr := data[:]
	smallestIndex, total := 0, len(data)

	for !sorted {
		if smallestIndex < total {
			smallest := arr[smallestIndex]
			iterations++
			for index := smallestIndex + 1; index < total; index++ {
				elem := arr[index]
				comparisons++
				if elem < smallest {
					arr[index] = smallest
					arr[smallestIndex] = elem
					smallest = elem
				}
			}
			smallestIndex++
		} else {
			sorted = true
		}
	}

	if verbose {
		log.Printf("Sorted data %v", arr)
	}
	return formatResult(arr, iterations, comparisons, total)
}

func stats(res Result) {
	if res.Error != nil {
		log.Println(res.Error)
	} else {
		log.Printf("%d operations through %d items with %d comparisons made", res.Iterations, res.Length, res.Comparisons)
	}
}

// Insertion sort
// takes an array of unsorted integers and sorts them
// returns the number of operations used
// sorts in ascending order, smallest on the left, largest right
func insertion(data []int, verbose bool) (Result, error) {
	if verbose {
		log.Println("Insertion Sort")
		log.Println("==========================")
		log.Printf("Unsorted data %v", data)
	}

	// obtain a slice
	// return early if we have 1 or less items
	total := len(data)
	if total <= 1 {
		return formatResult(data, 0, 0, total)
	}

	iterations, comparisons := 0, 0

	return formatResult(data, iterations, comparisons, total)
}

// Sort will run the sort
func Sort(sortType int, sampleData []int, verboseOpts ...bool) ([]int, error) {
	var verbose bool
	if len(verboseOpts) < 1 {
		verbose = false
	} else {
		verbose = verboseOpts[0]
	}
	var sortError error
	switch sortType {
	case Bubble:
		res, err := bubble(sampleData, verbose)
		if verbose {
			stats(res)
		}
		return res.Data, err
	case Selection:
		res, err := selection(sampleData, verbose)
		if verbose {
			stats(res)
		}
		return res.Data, err
	case Insertion:
		res, err := insertion(sampleData, verbose)
		if verbose {
			stats(res)
		}
		return res.Data, err
	default:
		sortError = errors.New("That sort method does not exist")
	}
	return nil, sortError
}
