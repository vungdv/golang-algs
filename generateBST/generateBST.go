package main

func generateTreesNoCache(n int) []*TreeNode {
	var buildTrees func(start, end int) []*TreeNode
	buildTrees = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}

		var allTrees []*TreeNode
		for i := start; i <= end; i++ {
			// for each root of val from start -> end,
			// root split start-end into two sub-trees
			// lefts < root < rights as they are BST.

			// build all possible sub-trees for start-(root-1)
			leftTrees := buildTrees(start, i-1)
			// build all possible sub-trees for (root+1) - end
			rightTrees := buildTrees(i+1, end)

			// composite all possible left-root-right
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					root := &TreeNode{Val: i}
					root.Left = left
					root.Right = right
					allTrees = append(allTrees, root)
				}
			}
		}
		return allTrees
	}
	return buildTrees(1, n)
}
