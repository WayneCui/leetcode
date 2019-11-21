import java.util.Map;

/*
 * @lc app=leetcode id=1189 lang=java
 *
 * [1189] Maximum Number of Balloons
 */

// @lc code=start
class Solution0 {
    public int maxNumberOfBalloons(String text) {
        Map<String, Integer> memo = new HashMap<>();
        for(int i = 0; i < text.length(); i++) {
            String key = String.valueOf(text.charAt(i));
            memo.put(key, memo.getOrDefault(key, 0) + 1);
        }

        String[] chars = new String[]{"b", "a", "l", "o", "n"};
        int count = Integer.MAX_VALUE;
        for (int i = 0; i < chars.length; i++) {
            if(chars[i] == "l" || chars[i] == "o") {
                count = Math.min(count, memo.getOrDefault(chars[i], 0) / 2);
            } else {
                count = Math.min(count, memo.getOrDefault(chars[i], 0));
            }
        }

        return count;
    }
}


class Solution {
    public int maxNumberOfBalloons(String text) {
        char[] dict = new char[26];
        for(char c: text.toCharArray()) {
            dict[c - 'a'] += 1;
        }

        int count = Integer.MAX_VALUE;
        count = Math.min(count, dict['b' - 'a']);
        count = Math.min(count, dict['a' - 'a']);
        count = Math.min(count, dict['l' - 'a'] / 2);
        count = Math.min(count, dict['o' - 'a'] / 2);
        count = Math.min(count, dict['n' - 'a']);

        return count;
    }
}
// @lc code=end

