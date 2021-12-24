package main

func inorderTraversal(root *TreeNode) []int {
	nodes := []int{}
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		nodes = append(nodes, root.Val)
		helper(root.Right)
	}
	helper(root)
	return nodes
}

func preorderTraversal(root *TreeNode) []int {
	nodes := []int{}
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		nodes = append(nodes, root.Val)
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)
	return nodes
}

//https://leetcode.com/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *TreeNode) []int {
	nodes := []int{}
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		helper(node.Right)
		nodes = append(nodes, node.Val)
	}
	helper(root)
	return nodes
}

func levelOrder(root *TreeNode) [][]int {
	var levelNodes [][]int
	if root == nil {
		return levelNodes
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nodes := []int{}
		nextq := []*TreeNode{}
		for _, curr := range q {
			if curr.Left != nil {
				nextq = append(nextq, curr.Left)
			}
			if curr.Right != nil {
				nextq = append(nextq, curr.Right)
			}
			nodes = append(nodes, curr.Val)
		}
		levelNodes = append(levelNodes, nodes)
		q = nextq
	}
	return levelNodes
}

func getLevelNodes(root *TreeNode) []int {
	var levelNodes []int
	if root == nil {
		return levelNodes
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextq := []*TreeNode{}
		for _, curr := range q {
			if curr.Left != nil {
				nextq = append(nextq, curr.Left)
			}
			if curr.Right != nil {
				nextq = append(nextq, curr.Right)
			}
			levelNodes = append(levelNodes, curr.Val)
		}
		q = nextq
	}
	return levelNodes
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*TreeNode{root}
	res := []float64{}

	for len(queue) != 0 {
		n, sum := len(queue), 0
		for i := 0; i < n; i++ {
			pop := queue[0]
			queue = queue[1:]
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
			sum += pop.Val
		}
		res = append(res, float64(sum)/float64(n))
	}
	return res
}
