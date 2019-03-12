/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.15%)
 * Total Accepted:    248.7K
 * Total Submissions: 773.6K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 * 
 * If the last word does not exist, return 0.
 * 
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 * 
 * Example:
 * 
 * Input: "Hello World"
 * Output: 5
 * 
 * 
 */
func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i + 1 < len(s) && string(s[i]) == " " && string(s[i+1]) != " " {
			return length
		}
		if string(s[i]) != " " {
			length++
		}
	}
	return length
}
