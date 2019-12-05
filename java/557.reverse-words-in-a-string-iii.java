/*
 * @lc app=leetcode id=557 lang=java
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
class Solution {
    public String reverseWords(String s) {
        byte[] bytes = (s + " ").getBytes();
        int prev = -1;
        int n = bytes.length;
        for (int i = 0; i < n; i++) {
            if(bytes[i] == ' '){
                int low = prev + 1;
                int high = i - 1;
                for(; low < high; low++, high--){
                    byte tmp = bytes[low];
                    bytes[low] = bytes[high];
                    bytes[high] = tmp;
                }
                prev = i;
            }
        }

        return new String(bytes, 0, bytes.length - 1);
    }
}
// @lc code=end

