/*
 * @lc app=leetcode id=17 lang=java
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
import java.util.*;

class Solution {
    private String[] dict = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};
    private int num;
    private String digits;

    public List<String> letterCombinations(String digits) {
        char[] dChars = digits.toCharArray();
        this.num = dChars.length;
        this.digits = digits;

        if(digits.isEmpty()) {
            return Collections.EMPTY_LIST;
        }

        List<String> result = new ArrayList<>();
        backtrace("", 0, result);
        return result;
    }

    public void backtrace(String soFar, int i, List<String> result) {
        if(i > num - 1) {
            result.add(soFar);
            return;
        }

        String letters = dict[Integer.parseInt(digits.charAt(i) + "")];
        if(letters.isEmpty()) {
            StringBuilder buf = new StringBuilder(soFar);
            backtrace(buf.toString(), i + 1, result);
        } else {
            char[] cs = letters.toCharArray();
            for(char c: cs) {
                StringBuilder buf = new StringBuilder(soFar);
                buf.append(c);
                backtrace(buf.toString(), i + 1, result);
            }
        }
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.letterCombinations("1"));
        System.out.println(s.letterCombinations("23"));
        System.out.println(s.letterCombinations("213"));
        System.out.println(s.letterCombinations("2345"));
    }
}

// @lc code=end

