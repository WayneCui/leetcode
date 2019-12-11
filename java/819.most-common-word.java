/*
 * @lc app=leetcode id=819 lang=java
 *
 * [819] Most Common Word
 */

// @lc code=start
import java.util.*;
class Solution {
    public String mostCommonWord(String paragraph, String[] banned) {
        String[] words = paragraph.replaceAll("\\W+", " ").toLowerCase().split("\\s+");
        Set<String> bannedSet = new HashSet<>(Arrays.asList(banned));

        Map<String, Integer> memo = new HashMap<String, Integer>();
        for(String word: words) {
            if(!bannedSet.contains(word)) {
                memo.put(word, memo.getOrDefault(word, 0) + 1);
            }
        }

        return Collections.max(memo.entrySet(), Map.Entry.comparingByValue()).getKey();
    }
}
// @lc code=end

