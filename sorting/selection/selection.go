package selection

import (
	"fmt"
	"time"
)

type Selection struct {
	Data interface{}
}

// New is an initialized for selection sort.
func New(data interface{}) *Selection {
	return &Selection{
		Data: data,
	}
}

// SortInteger is a function for sorting integer data type.
func (s *Selection) SortInteger() {
	newData := s.Data.([]int)
	t := time.Now()
	var temp int
	for i := 0; i < len(newData); i++ {
		min := i
		for j := i + 1; j < len(newData); j++ {
			if newData[j] < newData[min] {
				min = j
			}
		}
		if i != min {
			temp = newData[i]
			newData[i] = newData[min]
			newData[min] = temp

		}
	}
	fmt.Println("Selection : ", time.Since(t))
}
