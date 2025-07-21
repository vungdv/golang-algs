package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	trees := generateTreesNoCache(16)
	fmt.Println("Cache: ", CacheHitCount)
	fmt.Println("Total Trees: ", len(trees))
}

func (root *TreeNode) Println() {
	if root == nil {
		return
	}
	var visit func(node *TreeNode, t string)
	visit = func(node *TreeNode, t string) {
		if node == nil {
			return
		}
		//visit
		fmt.Printf("%s: %d, ", t, node.Val)

		visit(node.Left, "Left")
		visit(node.Right, "Right")
	}
	visit(root, "Root")
}
