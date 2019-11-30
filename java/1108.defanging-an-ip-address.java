/*
 * @lc app=leetcode id=1108 lang=java
 *
 * [1108] Defanging an IP Address
 */

// @lc code=start
class Solution {
    public String defangIPaddr(String address) {
        StringBuilder buf = new StringBuilder();
        for (int i = 0; i < address.length(); i++) {
            if(address.charAt(i) == '.') {
                buf.append('[');
                buf.append(address.charAt(i));
                buf.append(']');
            } else {
                buf.append(address.charAt(i));
            }
            
        }

        return buf.toString();
    }
}
// @lc code=end

