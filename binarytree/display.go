package main

import (
	"fmt"
	"strconv"
)

//------------------------------------------------------------------------------

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//------------------------------------------------------------------------------

func display(root *TreeNode) {
	treeString(root)
}

//------------------------------------------------------------------------------

func treeString(root *TreeNode) {
	str := toTreeString("", true, "", root)
	fmt.Println(str)
}

//------------------------------------------------------------------------------

func toTreeString(prefix string, top bool, str string, node *TreeNode) string {
	if node == nil {
		return ""
	}

	left := node.Left
	right := node.Right

	if right != nil {
		temp := path(top, ""+prefix, "│   ", "    ")
		str = toTreeString(temp, false, str, right)
	}

	str = path(top, str+prefix, "└──", "┌──")
	str = str + " " + strconv.Itoa(node.Val) + "\n"

	if left != nil {
		temp := path(top, ""+prefix, "    ", "│   ")
		str = toTreeString(temp, true, str, left)
	}

	return str
}

//------------------------------------------------------------------------------

func path(condition bool, str string, choice1 string, choice2 string) string {

	if condition {
		str += choice1
	} else {
		str += choice2
	}

	return str
}

func getTreeString(root *TreeNode) string {
	nodes := getLevelNodes(root)
	treeStr := "["
	for i, node := range nodes {
		if i < len(nodes)-1 {
			treeStr += fmt.Sprintf("%d,", node)
		} else {
			treeStr += fmt.Sprintf("%d]", node)
		}
	}
	return treeStr
}
