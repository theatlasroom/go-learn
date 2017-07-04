package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	fmt.Println("Starting")

	wg.Add(1)
	go func() {
		PAUSE := time.Duration(2 * time.Second)
		fmt.Println("........")
		time.Sleep(PAUSE)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Finished")
}
