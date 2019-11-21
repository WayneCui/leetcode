/*
 * @lc app=leetcode id=1221 lang=java
 *
 * [1221] Split a String in Balanced Strings
 */

// @lc code=start
import java.util.Stack;
class Solution {
    public int balancedStringSplit(String s) {
        char[] chars = s.toCharArray();
        int ls = 0, rs = 0;
        int count = 0;
        for(char c: chars) {
            if(ls > 0 && ls == rs) {
                count += 1;
                ls = rs = 0;
            }

            if(c == 'L') {
                ls += 1;
            } else {
                rs += 1;
            }
        }

        if(ls > 0 && ls == rs) {
            count += 1;
        }

        return count;
    }

    public int balancedStringSplit2(String s) {
        char[] chars = s.toCharArray();
        Stack<Character> stack = new Stack<>();
        char first = 0;
        int count = 0;

        for(char c: chars) {
            if(stack.isEmpty()) {
                first = c;
                count += 1;
            }

            if(c == first) {
                stack.push(c);
            } else {
                stack.pop();
            }
        }
        
        return count;
    }
}
// @lc code=end

