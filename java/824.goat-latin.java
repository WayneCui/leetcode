/*
 * @lc app=leetcode id=824 lang=java
 *
 * [824] Goat Latin
 */

// @lc code=start
import java.util.*;
class Solution {
    public String toGoatLatin(String S) {
        List<Character> list = Arrays.asList('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U');
        String[] tokens = S.split("\\s+");
        System.out.println(Arrays.toString(tokens));
        StringBuilder buf = new StringBuilder();
        for(int i = 0; i < tokens.length; i++) {
            String token = tokens[i];
            if(list.contains(token.charAt(0))) {
                buf.append(token);
            } else {
                buf.append(token.substring(1)).append(token.charAt(0));
            }

            buf.append("ma");
            for(int j = 0; j < i + 1; j++) {
                buf.append("a");
            }
            
            if(i < tokens.length - 1) {
                buf.append(" ");
            }
        }

        return buf.toString();
    }
}
// @lc code=end

