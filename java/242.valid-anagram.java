/*
 * @lc app=leetcode id=242 lang=java
 *
 * [242] Valid Anagram
 */

// @lc code=start
import java.util.HashMap;
import java.util.Map;
import java.util.Arrays;

class Solution {
    public boolean isAnagram(String s, String t) {
        if(s.length() != t.length()) {
            return false;
        }
        
        Map<Character, Integer> smap = new HashMap<>();
        char[] charsS = s.toCharArray();
        for(char c: charsS) {
            smap.put(c, (smap.getOrDefault(c, 0) + 1));
        }
        
        Map<Character, Integer> tmap = new HashMap<>();
        char[] charsT = t.toCharArray();
        for(char c: charsT) {
            tmap.put(c, (tmap.getOrDefault(c, 0) + 1));
        }
        
        if(smap.size() != tmap.size()) {
            return false;
        }
        
        for(Map.Entry<Character, Integer> entry: smap.entrySet()) {
            if(!tmap.containsKey(entry.getKey()) || !tmap.get(entry.getKey()).equals(entry.getValue())) {
                return false;
            }
        }
        
        return true;
    }

    public boolean isAnagram2(String s, String t) {
        int[] map = new int[26];
        for(char c: s.toCharArray()) {
            map[c - 'a']++;
        }
        
        for(char c: t.toCharArray()) {
            map[c - 'a']--;
        }
        
        for(int i: map) {
            if(i != 0) {
                return false;
            }
        }
        
        return true;
    }

    public boolean isAnagram3(String s, String t) {
        char[] c1 = s.toCharArray();
        char[] c2 = t.toCharArray();
        Arrays.sort(c1);
        Arrays.sort(c2);
        return Arrays.equals(c1, c2);
    }
}
// @lc code=end

