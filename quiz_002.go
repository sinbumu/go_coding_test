// 9. Palindrome Number

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Constraints:

// -231 <= x <= 231 - 1

// Follow up: Could you solve it without converting the integer to a string?

package main

import "strconv"

func isPalindrome(x int) bool {
	a := strconv.Itoa(x)
	if a[0:1] == "-" {
		return false
	}
	length := len(a)
	if length%2 == 0 {
		// fmt.Println(length)
		// fmt.Println(a[0:length/2])
		// fmt.Println(a[length/2:length])
		if a[0:length/2] == Reverse(a[length/2:length]) {
			return true
		}
	} else {
		// fmt.Println(length)
		// fmt.Println(a[0:length/2+1])
		// fmt.Println(a[length/2:length])
		if a[0:length/2+1] == Reverse(a[length/2:length]) {
			return true
		}
	}
	return false
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
