package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"reflect"
	"sync"
)

// Walk walks the tree t 
// Sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkTree(t, ch)
	defer close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {
	// Walk the left side of the tree
	if t.Left != nil {
		walkTree(t.Left, ch)
	}

	ch <- t.Value
	// Walk the right side of the tree
	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	arr1, arr2 := []int{}, []int{}

	// Create a channel for integers
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		Walk(t1, ch1)
		defer wg.Done()
	}()	
		
	for {
		i, ok := <- ch1
		if ok == false {
			break
		}
		arr1 = append(arr1, i)
	}

	go func(){ 
		Walk(t2, ch2)
		defer wg.Done()		
	}()

	for {	
		j, ok := <- ch2
		if ok == false {
			break
		}
		arr2 = append(arr2, j)
	}	

	wg.Wait()
	fmt.Printf("\narr1 %+v", arr1)
	fmt.Printf("\narr2 %+v", arr2)
	res := reflect.DeepEqual(arr1, arr2)
	return res
}

func main() {
	eq := Same(tree.New(1), tree.New(1))
	fmt.Printf("\nAre equal %b", eq)
}
