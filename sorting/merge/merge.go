package merge

import (
	"fmt"
	"math"
)

// Node is a struct to contain value from n of data, top, left, and right node.
// It used for creating tree or list.
type Node struct {
	Value []int
	Left  *Node
	Right *Node
}

// NewNode is a constructor.
func NewNode() *Node {
	return &Node{}
}

// Sort is a function to sort data.
// this function have a do split and merge data or we can call it
// "Divide and Conquer",
// what a cool name.
// Divide is to split data into tree until it only have single data. For detail check the Divide function.
// Conquer is to merge the data but this function also sort the value. for detail check the Conquer function.
func Sort(n []int) []int {
	node := &Node{
		Value: n,
	}
	Divide(node)
	Conquer(node)
	return node.Value
}

// Divide is a function to split the data into binary tree until the node value only have single data.
// Divide split the node value into two part, left and right
// this function do recursive until the node value only have single data.
// After the divide is done, then it will return a tree node.
func Divide(node *Node) {
	if node == nil {
		return
	}

	if len(node.Value) == 1 {
		return
	}

	node.Left = NewNode()
	node.Right = NewNode()

	// Why using ceil because len value not represented the array index.
	midIndex := int(math.Ceil(float64(len(node.Value)) / float64(2)))

	node.Left.Value = node.Value[0:midIndex]
	node.Right.Value = node.Value[midIndex:]

	Divide(node.Left)
	Divide(node.Right)

}

// Conquer is to merge and compare all separated node into single array in the top node.
func Conquer(node *Node) {
	if node.Left != nil {
		Conquer(node.Left)
	} else {
		return
	}
	node.Value = SortMergeLeftRightNodeValue(node)
	node.Left = nil

	if node.Right != nil {
		Conquer(node.Right)
	}
	node.Value = SortMergeLeftRightNodeValue(node)
	node.Right = nil
}

// SortMergeLeftRightNodeValue is to sort and merge two slice of value.
// This function must be called from the outer node or leaf, so every time this function is called
// the array is already sorted.
func SortMergeLeftRightNodeValue(node *Node) []int {
	var value, leftValue, rightValue []int

	if node.Right == nil {
		return node.Left.Value
	}

	if node.Left != nil {
		leftValue = node.Left.Value
	}

	if node.Right != nil {
		rightValue = node.Right.Value
	}

	var leftIndex, rightIndex int
	for {

		// condition if all left value already move, just move all right value.
		if leftIndex == len(leftValue) {
			value = append(value, rightValue[rightIndex:]...)
			return value
		}

		// condition if all right value already move, just move all left value.
		if rightIndex == len(rightValue) {
			value = append(value, leftValue[leftIndex:]...)
			return value
		}

		if leftValue[leftIndex] <= rightValue[rightIndex] {
			value = append(value, leftValue[leftIndex])
			leftIndex++
			continue
		}
		value = append(value, rightValue[rightIndex])
		rightIndex++
	}
}

// PrintNode is to print every node value from left to right g.
func PrintNode(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	PrintNode(node.Left)
	PrintNode(node.Right)
}
