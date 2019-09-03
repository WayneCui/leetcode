/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package main

import (
	"fmt"
	"container/list"
)
func isValid(s string) bool {
	stack := list.New()
	memo := make(map[rune]rune)
	memo[')'] = '('
	memo[']'] = '['
	memo['}'] = '{'
    for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.PushBack(c)
		} else {
			if stack.Len() > 0 {
				if memo[c] != stack.Remove(stack.Back()).(rune) {
					return false
				}
			} else {
				return false
			}
		}
	}

	if stack.Len() == 0 { return true } else { return false }
}

// func main() {
// 	s := "["
// 	fmt.Println(isValid(s))
// }
