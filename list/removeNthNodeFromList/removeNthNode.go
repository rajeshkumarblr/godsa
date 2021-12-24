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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return nil
	}

	nd := head
	for i := 0; i < n && nd != nil; i++ {
		nd = nd.Next
	}

	var prev *ListNode
	nthNode := head
	for ; nd != nil; nd = nd.Next {
		prev = nthNode
		nthNode = nthNode.Next
	}
	if prev != nil {
		prev.Next = nthNode.Next
	} else {
		head = head.Next
	}
	return head
}

func main() {
	nums1 := []int{1, 2}

	list1 := creatList(nums1)
	display(list1)

	list2 := removeNthFromEnd(list1, 2)
	display(list2)
}
