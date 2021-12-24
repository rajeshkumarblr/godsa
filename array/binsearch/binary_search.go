package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := (left + right) / 2
	for left <= right {
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return -1
}

func test(nums []int, target int) {
	ind := binarySearch(nums, target)
	if ind < 0 {
		fmt.Println("target: ", target, " not found.")
	} else {
		fmt.Println("index: ", ind, " val: ", nums[ind])
	}
}

func main() {
	nums := []int{10, 20, 30, 35, 36, 37, 38, 39, 40, 41, 42}
	for _, num := range nums {
		test(nums, num)
	}
	test(nums, 100)
	test(nums, 0)
	test(nums, 43)
}
