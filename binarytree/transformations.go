package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	root.Right = invertTree(left)
	root.Left = invertTree(right)
	return root
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	} else {
		root1.Val = root1.Val + root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
		return root1
	}
}

func runMergeTrees() {
	root1 := buildTree("1,3,2,5")
	display(root1)
	root2 := buildTree("2,1,3,null,4,null,7")
	display(root2)
	root3 := mergeTrees(root1, root2)
	display(root3)
}
