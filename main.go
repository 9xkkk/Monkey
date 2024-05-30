package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	q := []*TreeNode{root}
	p := []*TreeNode{subRoot}
	for len(q) > 0 && len(p) > 0 {
		node1 := q[0]
		node2 := p[0]
		q = q[1:]
		if node1.Val == node2.Val {
			queue := []*TreeNode{node1.Left, node1.Right}
			p = []*TreeNode{node2.Left, node2.Right}
			var x = false
			for len(queue) > 0 && len(p) > 0 {
				l1, l2 := queue[0], p[0]
				queue = queue[1:]
				p = p[1:]
				if l1 == nil && l2 == nil {
					queue, p = queue[1:], p[1:]
					continue
				}
				if l1 == nil || l2 == nil {
					x = true
					break
				}
				if l1.Val != l2.Val {
					x = true
					break
				}
				queue = append(queue, l1.Left)
				queue = append(queue, l1.Right)
				p = append(p, l2.Left)
				p = append(p, l2.Right)
			}
			if x {
				p = []*TreeNode{subRoot}
			}
		}
		if node1.Left != nil {
			q = append(q, node1.Left)
		}
		if node1.Right != nil {
			q = append(q, node1.Right)
		}
	}

	if len(p) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:  -9,
			Left: nil,
			Right: &TreeNode{
				Val:   -1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}

	subroot := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	ans := isSubtree(root, subroot)
	fmt.Println(ans)
}
