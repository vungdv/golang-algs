package main

// The idea is generate the tree of n nodes based on the previous one n-1, n-2, ...
// This lead to a very complicated logic as below
func generateTrees2(n int) []*TreeNode {
	insertRight := func(root TreeNode, v int) *TreeNode {
		curr := &root
		for curr.Right != nil {
			// copy right, assign curr -> copied node.
			// we need to do this because we are going to append to the right
			// by copy we avoid modified existing tree then it can be reusing.
			curr.Right = &TreeNode{
				Val:   curr.Right.Val,
				Left:  curr.Right.Left,
				Right: curr.Right.Right,
			}
			curr = curr.Right
		}

		curr.Right = &TreeNode{
			Val: v,
		}

		return &root
	}
	insertLeft := func(root TreeNode, v int) *TreeNode {
		curr := &root
		for curr.Left != nil {
			// copy Left, assign curr -> copied node.
			// we need to do this because we are going to append to the Left
			// by copy we avoid modified existing tree then it can be reusing.
			curr.Left = &TreeNode{
				Val:   curr.Left.Val,
				Left:  curr.Left.Left,
				Right: curr.Left.Right,
			}
			curr = curr.Left
		}

		curr.Left = &TreeNode{
			Val: v,
		}

		return &root
	}

	joinRightValueLeft := func(root TreeNode, v int, left *TreeNode) *TreeNode {
		curr := &root
		for curr.Right != nil {
			// copy Right, assign curr -> copied node.
			// we need to do this because we are going to append to the Right
			// by copy we avoid modified existing tree then it can be reusing.
			curr.Right = &TreeNode{
				Val:   curr.Right.Val,
				Left:  curr.Right.Left,
				Right: curr.Right.Right,
			}
			curr = curr.Right
		}

		curr.Right = &TreeNode{
			Val:  v,
			Left: left,
		}

		return &root
	}

	joinLeftValueRight := func(root TreeNode, v int, right *TreeNode) *TreeNode {
		curr := &root
		for curr.Left != nil {
			// copy right, assign curr -> copied node.
			// we need to do this because we are going to append to the right
			// by copy we avoid modified existing tree then it can be reusing.
			curr.Left = &TreeNode{
				Val:   curr.Left.Val,
				Left:  curr.Left.Left,
				Right: curr.Left.Right,
			}
			curr = curr.Left
		}

		curr.Left = &TreeNode{
			Val:   v,
			Right: right,
		}

		return &root
	}

	genCache := make(map[int][]*TreeNode)
	type IntPair struct {
		V     int
		Count int
	}

	genCacheReverse := make(map[IntPair][]*TreeNode)

	var generate func(m int) []*TreeNode
	var generateReverse func(m, count int) []*TreeNode
	generateReverse = func(m, count int) []*TreeNode {
		if result, exist := genCacheReverse[IntPair{V: m, Count: count}]; exist {
			return result
		}
		// m is gurantee >= 2
		if m < 2 || count <= 0 {
			return nil
		} else if count == 1 {
			return []*TreeNode{
				&TreeNode{
					Val: m,
				},
			}
		} else if count == 2 {
			return []*TreeNode{
				&TreeNode{
					Val: m,
					Left: &TreeNode{
						Val: m - 1,
					},
				},
				&TreeNode{
					Val: m - 1,
					Right: &TreeNode{
						Val: m,
					},
				},
			}
		}
		//f(n) =
		// root(m-count) -(right) f(m, count-1)
		// f(m, count-1) - left (m-count)

		// f(m, count-2) left (m-count) right f(m-2, 1)
		// f(m, count-3)  -   (m-count)   -   f(m-3, 2)
		// f(m, count-4)  -   (m-count)   -   f(m-4, 3)
		// f(m, count-5)  -   (m-count)   -   f(m-5, 4)
		// f(m, count-6)  -   (m-count)   -   f(m-6, 5)

		result := []*TreeNode{}
		preResult := generateReverse(m, count-1)
		for _, r := range preResult {
			// m as root
			result = append(result, &TreeNode{Val: m - count + 1, Right: r})
			// m as right
			result = append(result, insertLeft(*r, m-count+1))
		}

		for i := 3; i <= count; i++ {
			f1 := generateReverse(m, i-2)
			f2 := generateReverse(m-i+2, count-i+1)
			for _, rf1 := range f1 {
				for _, rr := range f2 {
					// we copy rf1 because it will append right,
					// that will avoid modifing existing tree.
					result = append(result, joinLeftValueRight(*rf1, m-count+1, rr))
				}
			}
		}

		genCacheReverse[IntPair{V: m, Count: count}] = result
		return result
	}

	generate = func(m int) []*TreeNode {
		if result, exit := genCache[m]; exit {
			return result
		}

		if m == 1 {
			return []*TreeNode{
				&TreeNode{
					Val: 1,
				},
			}
		}

		if m == 2 {
			return []*TreeNode{
				&TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				&TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
			}
		}
		//f(n) = root -(left) f(n-1), f(n-1) - (left) n
		// f(n-2) - right n - left f(n-1)
		// f(n-3) - right n - left f(n-1, n-2)
		// f(n-4) - right n - left f(n-1, n-2, n-3)
		// ...
		// f(2) - right n - left f(n-1, n-2, ... 3)
		// f(1) - right n - left f(n-1, n-2, ... 2)
		result := []*TreeNode{}
		preResult := generate(m - 1)
		for _, r := range preResult {
			// m as root
			result = append(result, &TreeNode{Val: m, Left: r})
			// m as right
			result = append(result, insertRight(*r, m))
		}

		for i := m - 2; i > 0; i-- {
			f1 := generate(i)
			f2 := generateReverse(m-2+1, m-i-1)
			for _, rf1 := range f1 {
				for _, rr := range f2 {
					// we copy rf1 because it will append right,
					// that will avoid modifing existing tree.
					result = append(result, joinRightValueLeft(*rf1, m, rr))
				}
			}
		}

		genCache[m] = result
		return result
	}

	return generate(n)
}
