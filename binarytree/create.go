package main

import (
	"strconv"
	"strings"
)

func addNode(node *TreeNode, isLeft bool, numPtr *int) *TreeNode {
	var newNode *TreeNode = nil
	if numPtr != nil {
		newNode = &TreeNode{
			Val:   *numPtr,
			Left:  nil,
			Right: nil,
		}
	}
	if isLeft {
		node.Left = newNode
	} else {
		node.Right = newNode
	}
	return newNode
}

func buildTree(inputStr string) *TreeNode {
	inputs := strings.Split(inputStr, ",")
	nums := []*int{}
	for _, input := range inputs {
		if input == "null" {
			nums = append(nums, nil)
		} else {
			val, err := strconv.Atoi(input)
			if err != nil {
				continue
			}
			nums = append(nums, &val)
		}
	}
	if len(nums) == 0 {
		return nil
	}

	q := []*TreeNode{}
	i := 0
	root := &TreeNode{
		Val:   *nums[0],
		Left:  nil,
		Right: nil,
	}
	q = append(q, root)
	i++
	for len(q) > 0 && i < len(nums) {
		tmp := nums[i]
		node := q[0]
		if node != nil {
			newNode := addNode(node, true, tmp)
			q = append(q, newNode)
			i++
			if i >= len(nums) {
				break
			}
		}
		tmp = nums[i]
		if node != nil {
			newNode := addNode(node, false, tmp)
			q = append(q, newNode)
			i++
			if i >= len(nums) {
				break
			}
		}
		q = q[1:]
	}
	return root
}
