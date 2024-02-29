package Leetcode

import (
	"sort"
)

// twoSum Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
func twoSum(input []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(input); i++ {
		_, ok := m[input[i]]
		if !ok {
			m[target-input[i]] = i
		} else {
			return []int{m[input[i]], i}
		}
	}
	return nil
}

// threeSum Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets
func threeSum(input []int) [][]int {
	sort.Ints(input)
	var result [][]int
	for i := 0; i < len(input)-2; i++ {
		if i > 0 && input[i] == input[i-1] {
			continue
		}

		left, right := i+1, len(input)-1
		for left < right {
			sum := input[i] + input[left] + input[right]
			if sum == 0 {
				result = append(result, []int{input[i], input[left], input[right]})
				for left < right && input[left] == input[left+1] {
					left++
				}
				for left < right && input[right] == input[right-1] {
					right++
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// maxSubArray 53. Maximum Subarray
// Given an integer array nums, find the
// subarray
// with the largest sum, and return its sum.
func maxSubArray(nums []int) int {
	//currentSum := nums[0]
	//
	//for i, _ := range nums {
	//	tempSum := 0
	//	for j := i; j < len(nums); j++ {
	//		tempSum += nums[j]
	//		if tempSum > currentSum {
	//			currentSum = tempSum
	//		}
	//	}
	//}

	maxSum, currentSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < currentSum+nums[i] {
			currentSum = currentSum + nums[i]
		} else {
			currentSum = nums[i]
		}
		if maxSum < currentSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
