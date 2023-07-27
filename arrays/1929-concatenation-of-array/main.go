package main

import (
	"fmt"
)
// Given an integer array nums of length n, you want to create an array ans of length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).

// Specifically, ans is the concatenation of two nums arrays.

// Return the array ans.

// initial solution
func getConcatenation(nums []int) []int {	
	nums = append(nums,nums...)
	return nums
}

func main() {
	input := []int{0,1,2}
	output := getConcatenation(input)
	fmt.Println(output)
}
