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

    public String multiply2(String num1, String num2) {
        if("0".equals(num1) || "0".equals(num2)) {
            return "0";
        }

        int len1 = num1.length();
        int len2 = num2.length();
        int lenP = len1 + len2;
        char[] chars = new char[lenP];
        char[] chars1 = num1.toCharArray();
        char[] chars2 = num2.toCharArray();

        for(int i = len1 - 1; i >= 0; i--) {
            int c = chars1[i] - '0';
            for(int j = len2 - 1; j >= 0; j--) {
                chars[i + j + 1] += c * (chars2[j] - '0');
            }
        }

        for(int k = lenP - 1; k >= 0; k--) {
            if(chars[k] > 9) {
                chars[k - 1] += chars[k] / 10;
                chars[k] %= 10;
            }

            chars[k] += '0';
        }

        if(chars[0] == '0') {
            return new String(Arrays.copyOfRange(chars, 1, lenP));
        } else {
            return new String(chars);
        }
    }
}
// @lc code=end

