package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	L := dfs(root.Left, ans)
	R := dfs(root.Right, ans)
	*ans = Max(*ans, L+R)

	return Max(L, R) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	dia := 0
	dfs(root, &dia)
	return dia
}
