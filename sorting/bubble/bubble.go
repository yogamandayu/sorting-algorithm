package bubble

import (
	"fmt"
	"time"
)

type Bubble struct {
	Data interface{}
}

// New is an initialized for bubble sort.
func New(data interface{}) *Bubble {
	return &Bubble{
		Data: data,
	}
}

// SortInteger is a function for sorting integer data type.
func (b *Bubble) SortInteger() {
	var temp int
	t := time.Now()
	newData := b.Data.([]int)
	for i := 0; i < len(b.Data.([]int)); i++ {
		for j := 1; j < len(b.Data.([]int))-i; j++ {
			if newData[j-1] > newData[j] {
				temp = newData[j-1]
				newData[j-1] = newData[j]
				newData[j] = temp
			}
		}
	}
	fmt.Println("Bubble : ", time.Since(t))
}
