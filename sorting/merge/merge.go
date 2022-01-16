package merge

import (
	"fmt"
	"math"
	"time"
)

type Merge struct {
	Data interface{}
}

// New is an initialized for merge sort.
func New(data interface{}) *Merge {
	return &Merge{
		Data: data,
	}
}

type List struct {
	Top   *List
	Data  []int
	Left  *List
	Right *List
}

// SortIntegerWithList is a function for sorting integer data type with list structure.
func (s *Merge) SortIntegerWithList() {
	newData := s.Data.([]int)
	l := List{
		Data: newData,
	}
	t := time.Now()
	//printList(&l)
	s.Split(&l)
	s.Merge(&l)
	//fmt.Println(l.Data)
	//printList(&l)
	fmt.Println("Merge : ", time.Since(t))
}

// Split is a function to create a tree that filled with array and split them left and right until single array value.
func (s *Merge) Split(l *List) {
	if len(l.Data) > 1 {
		splitIndex := int(math.Floor(float64(len(l.Data) / 2)))
		l.Left = &List{
			Top:  l,
			Data: l.Data[0:splitIndex],
		}
		s.Split(l.Left)

		l.Right = &List{
			Top:  l,
			Data: l.Data[splitIndex:(len(l.Data))],
		}
		s.Split(l.Right)
	}
}

func (s *Merge) Merge(l *List) {
	if l.Left != nil { // if left is not nil, then right also not nil
		s.Merge(l.Left)
		s.Merge(l.Right)

		var i, j int
		l.Data = nil
		for {
			left := l.Left.Data
			right := l.Right.Data

			if i < len(left) {
				if left[i] <= right[j] {
					l.Data = append(l.Data, left[i])
					i++
				} else if right[j] <= left[i] {
					l.Data = append(l.Data, right[j])
					j++
				}
			}

			if i >= len(left) && j >= len(right) {
				break
			}

			if i >= len(left) {
				l.Data = append(l.Data, right[j:(len(right))]...)
				break
			}

			if j >= len(right) {
				l.Data = append(l.Data, left[i:(len(left))]...)
				break
			}
		}
		l.Left = nil  //reduce memory usage
		l.Right = nil //reduce memory usage
	}
}

func printList(l *List) {
	if l.Left != nil { // if left is not nil, then right also not nil
		printList(l.Left)
		printList(l.Right)
	}

	if l.Left == nil {
		fmt.Println(l.Data) // if left is nil, then right also nil
	}
}
