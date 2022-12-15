package merge_test

import (
	"Sorting/sorting/merge"
	"math/rand"

	"testing"
	"time"
)

func TestSort(t *testing.T) {
	t.Run("Run merge sort test with n is an array of shuffle number, expected no error and returning ordered array", func(t *testing.T) {
		arr := []int{176, 22, 67, 45, 33, 90, 77, 326, 78, 65, 48, 27}
		orderedArr := merge.Sort(arr)
		prevNumber := orderedArr[0]
		for _, currentNumber := range orderedArr {
			if currentNumber < prevNumber {
				t.Fail()
			}
			prevNumber = currentNumber
		}
	})

	t.Run("Run merge sort test with generated random n slice of int, expected no error and returning ordered array", func(t *testing.T) {
		rand.Seed(time.Now().Unix())
		arr := rand.Perm(1000000)
		orderedArr := merge.Sort(arr)
		prevNumber := orderedArr[0]
		for _, currentNumber := range orderedArr {
			if currentNumber < prevNumber {
				t.Fail()
			}
			prevNumber = currentNumber
		}
	})
}

func TestPrint(t *testing.T) {
	arr := []int{176, 22, 67, 45, 33, 90, 77, 326, 78, 65, 48, 27}

	node := &merge.Node{
		Value: arr,
	}
	merge.Divide(node)
	merge.Conquer(node)
	merge.PrintNode(node)
}

func TestDivide(t *testing.T) {
	t.Run("Run divide function with expected result is every node leaf value must be only have one single data", func(t *testing.T) {
		arr := []int{176, 22, 67, 45, 33, 90, 77, 326, 78, 65, 48, 27}

		node := &merge.Node{
			Value: arr,
		}

		merge.Divide(node)
		ok := true

		CheckNodeLeaf(node, ok)
		if !ok {
			t.Fail()
		}
	})
	t.Run("Run divide function with nil data, expected no error", func(t *testing.T) {
		merge.Divide(nil)
	})
}

func TestSortAndMerge(t *testing.T) {
	t.Run("Run sort and merge function with filled left and right node, expected return ordered value", func(t *testing.T) {
		node := &merge.Node{}

		node.Left = &merge.Node{
			Value: []int{2, 3, 7, 9},
		}

		node.Right = &merge.Node{
			Value: []int{4, 5, 7, 8},
		}
		value := merge.SortMergeLeftRightNodeValue(node)

		if len(value) != len(node.Left.Value)+len(node.Right.Value) {
			t.Fail()
		}

		prevNumber := value[0]
		for _, currentNumber := range value {
			if currentNumber < prevNumber {
				t.Fail()
			}
		}
	})

	t.Run("Run sort and merge function with filled left only, expected return ordered value", func(t *testing.T) {
		node := &merge.Node{}

		node.Left = &merge.Node{
			Value: []int{2, 3, 7, 9},
		}

		value := merge.SortMergeLeftRightNodeValue(node)
		if len(value) != len(node.Left.Value) {
			t.Fail()
		}

		prevNumber := value[0]
		for _, currentNumber := range value {
			if currentNumber < prevNumber {
				t.Fail()
			}
		}
	})

	t.Run("Run sort and merge function with filled right only, expected return ordered value", func(t *testing.T) {
		node := &merge.Node{}

		node.Right = &merge.Node{
			Value: []int{2, 3, 7, 9},
		}

		value := merge.SortMergeLeftRightNodeValue(node)
		if len(value) != len(node.Right.Value) {
			t.Fail()
		}

		prevNumber := value[0]
		for _, currentNumber := range value {
			if currentNumber < prevNumber {
				t.Fail()
			}
		}
	})
}

func CheckNodeLeaf(node *merge.Node, ok bool) {
	if node == nil {
		return
	}
	if node.Left == nil && len(node.Value) != 1 {
		ok = false
		return
	}
	CheckNodeLeaf(node.Left, ok)
	CheckNodeLeaf(node.Right, ok)
}
