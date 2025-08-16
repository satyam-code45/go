package main

import (
	"fmt"
	"slices"
)

func main() {
	//nil init
	// var nums[]int
	// fmt.Println(nums)

	var nums = make([]int, 0, 4)
	fmt.Println(nums)
	nums = append(nums, 11)
	//capacity -> max no of elements that can be fit.
	fmt.Println(cap(nums))
	fmt.Println(len(nums))
	nums = append(nums, 45)
	fmt.Println(nums)

	//capacity -> max no of elements that can be fit.
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

	//copy
	var nums2 = make([]int, len(nums))

	copy(nums2,nums)
	fmt.Println(nums, nums2)

	//slice operator
	fmt.Println(nums[:1])
	fmt.Println(nums[1:2])

	//comparing the slices
	fmt.Println(slices.Equal(nums,nums2))

}
