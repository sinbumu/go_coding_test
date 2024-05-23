/*
3151. Special Array I

An array is considered special if every pair of its adjacent elements contains two numbers with different parity.

You are given an array of integers nums. Return true if nums is a special array, otherwise, return false.

Example 1:

Input: nums = [1]

Output: true

Explanation:

There is only one element. So the answer is true.

Example 2:

Input: nums = [2,1,4]

Output: true

Explanation:

There is only two pairs: (2,1) and (1,4), and both of them contain numbers with different parity. So the answer is true.

Example 3:

Input: nums = [4,3,1,6]

Output: false

Explanation:

nums[1] and nums[2] are both odd. So the answer is false.

Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/
package main

func isArraySpecial(nums []int) bool {
	// 돌면서 직전값 캐싱해두고 대조하면 될듯?
	tempNum := 0
	//loop
	for index, num := range nums {
		if index == 0 {
			tempNum = num
			continue
		}
		if tempNum%2 == 0 && num%2 == 0 {
			return false
		} else if tempNum%2 != 0 && num%2 != 0 {
			return false
		}
		tempNum = num
	}
	return true
}
