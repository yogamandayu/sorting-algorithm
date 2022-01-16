package main

import (
	"Sorting/sorting/bubble"
	"Sorting/sorting/insertion"
	"Sorting/sorting/merge"
	"Sorting/sorting/selection"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	data := rand.Perm(100000)
	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		bubble.New(data).SortInteger()
		wg.Done()
	}()

	go func() {
		insertion.New(data).SortInteger()
		wg.Done()
	}()

	go func() {
		selection.New(data).SortInteger()
		wg.Done()
	}()

	go func() {
		merge.New(data).SortIntegerWithList()
		wg.Done()
	}()
	wg.Wait()
}
