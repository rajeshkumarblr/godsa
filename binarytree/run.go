package main

import "fmt"

func runAverageOfLevels() {
	//fmt.Println(averageOfLevels(root))
}

func runLCA() {
	root := buildTree("6,2,8,0,4,7,9,null,null,3,5")
	display(root)
	p := root.Left.Right.Left
	q := root.Left.Right.Right
	node := lowestCommonAncestor(root, p, q)
	fmt.Println("lca: ", node.Val)
}

func runMaxPathSum() {
	/*root := buildTree("2,-1")
	display(root)
	maxsum := maxPathSum(root)
	fmt.Println("max:", maxsum)
	fmt.Println()

	root = buildTree("-3")
	display(root)
	maxsum = maxPathSum(root)
	fmt.Println("max:", maxsum)
	fmt.Println()

	root = buildTree("-10,9,20,null,null,15,7")
	display(root)
	maxsum = maxPathSum(root)
	fmt.Println("max:", maxsum)
	fmt.Println()

	root = buildTree("1,-2,3")
	display(root)
	maxsum = maxPathSum(root)
	fmt.Println("max:", maxsum)
	fmt.Println()

	root = buildTree("5,4,8,11,null,13,4,7,2,null,null,null,1")
	display(root)
	maxsum = maxPathSum(root)
	fmt.Println("max:", maxsum)
	fmt.Println()*/
}

func runMaxDepth() {
	root := buildTree("5,4,8,11,null,13,4,7,2,null,null,null,1")
	display(root)
	depth := maxDepth(root)
	fmt.Println("depth:", depth)
}

func runInOrderTraversal() {
	root := buildTree("5,4,8,11,null,13,4,7,2,null,null,null,1")
	nodes := inorderTraversal(root)
	fmt.Println("nodes:", nodes)
}

func runIsSymmetric() {
	//root := buildTree("5,1,4,null,null,3,6")
	//root := buildTree("2,1,3")
	root := buildTree("32,26,47,19,null,null,56,null,27")
	display(root)
	flag := isValidBST(root)
	fmt.Println("isValidBST:", flag)
}

func runIsValidBST() {
	//root := buildTree("5,1,4,null,null,3,6")
	//root := buildTree("2,1,3")
	root := buildTree("32,26,47,19,null,null,56,null,27")
	display(root)
	flag := isValidBST(root)
	fmt.Println("isValidBST:", flag)
}

func runIsBalanced() {
	root := buildTree("1,2,2,3,null,null,3,4,null,null,4")
	display(root)
	flag := isBalanced(root)
	fmt.Println("isBalanced:", flag)
}
