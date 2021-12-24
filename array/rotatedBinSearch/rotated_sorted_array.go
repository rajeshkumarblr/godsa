package main

import (
	"fmt"
)

/*func findPivotElement(nums []int, left int, right int) int {
	if left == right {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] > nums[mid+1] {
		return mid
	} else if nums[left] < nums[mid] {
		return findPivotElement(nums, mid+1, right)
	} else {
		return findPivotElement(nums, left, mid)
	}
}*/

func findPivotElement(nums []int, left int, right int) int {
	mid := (left + right) / 2
	for nums[mid] < nums[mid+1] {
		if left == right {
			return -1
		}
		if nums[left] < nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return mid
}

func search(nums []int, target int) int {

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

func test(nums []int, target int) {
	ind := search(nums, target)
	if ind < 0 {
		fmt.Println("target: ", target, " not found.")
	} else {
		fmt.Println("index: ", ind, " val: ", nums[ind])
	}
}

func main() {
	nums := []int{45, 50, 60, 70, 80, 81, 10, 20, 30, 35, 36, 37, 38, 39, 40, 41, 42}
	for _, num := range nums {
		test(nums, num)
	}
	test(nums, 100)
	test(nums, 0)
	test(nums, 43)
	nums = []int{3, 1}
	for _, num := range nums {
		test(nums, num)
	}
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	test(nums, 0)
	nums = []int{1, 3}
	for _, num := range nums {
		test(nums, num)
	}
}
