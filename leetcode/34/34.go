package main

import (
	"sort"
)

func searchRange(nums []int, target int) []int {
	begin := -1
	end := -1

	if len(nums) > 0 {
		begin = sort.SearchInts(nums, target)
		if begin < len(nums) && nums[begin] == target {
			end = sort.SearchInts(nums, target+1) - 1
		} else {
			begin = -1
		}
	}

	return []int{begin, end}
}
