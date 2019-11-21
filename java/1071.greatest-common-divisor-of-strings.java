/*
 * @lc app=leetcode id=1071 lang=java
 *
 * [1071] Greatest Common Divisor of Strings
 */

// @lc code=start
class Solution {
    public String gcdOfStrings(String str1, String str2) {
        int len1 = str1.length();
        int len2 = str2.length();

        while(true) {
            if(len1 > len2) {
                if(str1.startsWith(str2)) {
                    str1 = str1.substring(len2);
                    len1 = str1.length();
                } else {
                    return "";
                }
            } else if(len1 < len2) {
                if(str2.startsWith(str1)) {
                    str2 = str2.substring(len1);
                    len2 = str2.length();
                } else {
                    return "";
                }
            } else {
                return str1.equals(str2) ? str1 : "";
            }
        }
    }
}
// @lc code=end

