package main

import (
	"Sorting/sorting/merge"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	data := rand.Perm(2000)

	//bubble.New(data).SortInteger()
	//
	//insertion.New(data).SortInteger()
	//
	//selection.New(data).SortInteger()

	merge.New(data).SortIntegerWithList()
}
