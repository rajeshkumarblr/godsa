package main

import "math"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

//LC: https://leetcode.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	var isMirror func(p *TreeNode, q *TreeNode) bool
	isMirror = func(p *TreeNode, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}
		return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
	}
	return isMirror(root, root)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	curr := root
	for {
		switch {
		case curr.Val < p.Val && curr.Val < q.Val:
			curr = curr.Right
		case curr.Val > p.Val && curr.Val > q.Val:
			curr = curr.Left
		default:
			return curr
		}
	}
}

func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var dfs func(root *TreeNode, min, max int) bool
	dfs = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val < min || root.Val > max {
			return false
		}

		return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
	}

	return dfs(root, math.MinInt32, math.MaxInt32)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, isLeftBalanced := checkBalance(root.Left)
	rightDepth, isRightBalanced := checkBalance(root.Right)
	balanced := isLeftBalanced && isRightBalanced
	balanced = balanced && abs(leftDepth-rightDepth) <= 1
	return max(leftDepth, rightDepth) + 1, balanced
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, result := checkBalance(root)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}

}
