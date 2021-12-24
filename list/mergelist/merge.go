package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func display(list *ListNode) {
	for ; list != nil; list = list.Next {
		fmt.Printf("%d ", list.Val)
	}
	fmt.Println("")
}

var list3 *ListNode
var prev *ListNode

func addNode(tmp *ListNode) {
	nd := &ListNode{
		Val:  tmp.Val,
		Next: nil,
	}
	if prev == nil {
		list3 = nd
	} else {
		prev.Next = nd
	}
	prev = nd
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	node1 := list1
	node2 := list2

	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			addNode(node1)
			node1 = node1.Next
		} else {
			addNode(node2)
			node2 = node2.Next
		}
	}

	if node1 != nil {
		for node1 != nil {
			addNode(node1)
			node1 = node1.Next
		}
	}

	if node2 != nil {
		for node2 != nil {
			addNode(node2)
			node2 = node2.Next
		}
	}

	return list3
}

func creatList(nums []int) *ListNode {
	cnt := len(nums)
	if cnt == 0 {
		return nil
	}
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	prev := head
	for i := 1; i < cnt; i++ {
		nd := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		prev.Next = nd
		prev = nd
	}
	return head
}

func main() {
	nums1 := []int{0, 2, 4, 6, 8}
	nums2 := []int{1, 3, 5, 7, 9}

	list1 := creatList(nums1)
	display(list1)

	list2 := creatList(nums2)
	display(list2)

	list3 := mergeTwoLists(list1, list2)
	display(list3)
}
