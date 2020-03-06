/*
 * @lc app=leetcode id=43 lang=java
 *
 * [43] Multiply Strings
 */

// @lc code=start
import java.util.*;
class Solution {
    public String multiply(String num1, String num2) {

        if("0".equals(num1) || "0".equals(num2)) {
            return "0";
        }

        int len1 = num1.length();
        int len2 = num2.length();
        int lenP = len1 + len2;

        StringBuilder buf = new StringBuilder(paddingZero(lenP));
        for(int i = len2 - 1; i >= 0; i--) {
            int carry = 0;
            int bufIndex = lenP - (len2 - i);

            for(int j = len1 - 1; j >= 0; j--) {
                int product = (num2.charAt(i) - '0') * (num1.charAt(j) - '0') + carry;
                int sum = product + (buf.charAt(bufIndex) - '0');
                buf.setCharAt(bufIndex, (char)(sum % 10 + '0'));
                carry = sum / 10;
                bufIndex--;
            }

            buf.setCharAt(bufIndex, (char)(carry + '0'));
        }

        return buf.charAt(0) == '0' ? buf.substring(1) : buf.toString();
    }

    private String paddingZero(int howMany) {
        char[] chars = new char[howMany];
        Arrays.fill(chars, '0');
        return new String(chars);
    }
}
// @lc code=end

