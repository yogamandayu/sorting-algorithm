package insertion

import (
	"fmt"
	"time"
)

type Insertion struct {
	Data interface{}
}

// New is an initialized for insertion sort.
func New(data interface{}) *Insertion {
	return &Insertion{
		Data: data,
	}
}

// SortInteger is a function for sorting integer data type.
func (is *Insertion) SortInteger() {
	newData := is.Data.([]int)
	t := time.Now()
	i := 1
	for {
		if !(i < len(newData)) {
			break
		}
		x := newData[i]
		j := i - 1
		for {
			if !((j >= 0) && (newData[j] > x)) {
				break
			}
			newData[j+1] = newData[j]
			j--
		}
		newData[j+1] = x
		i++
	}
	fmt.Println("Insertion : ", time.Since(t))
}
