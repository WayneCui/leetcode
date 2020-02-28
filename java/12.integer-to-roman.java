/*
 * @lc app=leetcode id=12 lang=java
 *
 * [12] Integer to Roman
 */

// @lc code=start

import java.util.*;

class Solution {
    public String intToRoman(int num) {
        char[] symbols = {'I', 'V', 'X', 'L', 'C', 'D', 'M'};
        int remainder = 0;

        int i = 0;
        StringBuilder buf = new StringBuilder();
        while(num > 0) {
            remainder = num % 10;
            num = num / 10;

            if(remainder < 4) {
                buf.append(repeatChar(symbols[i], remainder));
            } else if(remainder == 4) {
                buf.append(new String(new char[]{symbols[i + 1], symbols[i]}));
            } else if(remainder < 9) {
                buf.append(repeatChar(symbols[i], remainder - 5));
                buf.append(new String(new char[] {symbols[i + 1]}));
            } else if(remainder == 9) {
                buf.append(new String(new char[]{symbols[i + 2], symbols[i]}));
            }

            i += 2;
        }

        return buf.reverse().toString();
    }

    public static String repeatChar(char c, int times) {
        char[] chars = new char[times];
        Arrays.fill(chars, c);
        return new String(chars);
    }
}
// @lc code=end

