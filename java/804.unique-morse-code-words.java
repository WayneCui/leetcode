/*
 * @lc app=leetcode id=804 lang=java
 *
 * [804] Unique Morse Code Words
 */

// @lc code=start
import java.util.HashMap;
import java.util.Map;
class Solution {
    public int uniqueMorseRepresentations(String[] words) {
        String[] dict = {".-","-...","-.-.","-..",".","..-.","--.",
            "....","..",".---","-.-",".-..","--","-.",
            "---",".--.","--.-",".-.","...","-",
            "..-","...-",".--","-..-","-.--","--.."};
        Map<String, Integer> map = new HashMap<>();

        for(String word: words) {
            StringBuilder buf = new StringBuilder();
            for(char c: word.toCharArray()) {
                buf.append(dict[c - 'a']);
            }

            map.put(buf.toString(), 1);
        }
        return map.size();
    }
}
// @lc code=end

