package main

import (
	"fmt"
	"math"
)

func main() {

	// valid pre-order BST:
	//	root(3)
	// 				-> right(9)
	// 			-> left(8)
	// 		->left(5)
	// 			right(6)
	fmt.Println("Valid pre-order BST:", math.MinInt64)
	fmt.Println("Valid pre-order BST:", math.MaxInt64)
	test1 := []int{3, 9, 8, 5, 6}
	test2 := []int{3, 9, 8, 5, 10} // false sub-left tree of 8 shouldn't greater than 8
	test3 := []int{3, 9, 2, 5, 6}  // false, sub-right tree of 3 shouldn't less than 3
	const message = "Is valid pre-order BST:"

	fmt.Println(message, validatePreOrderBST(test1))
	fmt.Println(message, validatePreOrderBST(test2))
	fmt.Println(message, validatePreOrderBST(test3))

}

func validatePreOrderBST(arr []int) bool {
	if len(arr) < 2 {
		return true
	}

	if arr[0] < arr[1] {
		return validateRight(arr[0], 1, arr)
	} else {
		return validateLeft(arr[0], 1, arr)
	}
}

func validateLeft(k, index int, arr []int) bool {
	if index == len(arr)-1 {
		return true
	}

	if arr[index+1] > k {
		return false
	}

	if arr[index+1] > arr[index] {
		return validateRight(arr[index], index+1, arr)
	} else {
		return validateLeft(arr[index], index+1, arr)
	}
}

func validateRight(k, index int, arr []int) bool {
	if index == len(arr)-1 {
		return true
	}

	if arr[index+1] < k {
		return false
	}

	if arr[index+1] < arr[index] {
		return validateLeft(arr[index], index+1, arr)
	} else {
		return validateRight(arr[index], index+1, arr)
	}
}
