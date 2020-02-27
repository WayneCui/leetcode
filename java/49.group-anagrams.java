/*
 * @lc app=leetcode id=49 lang=java
 *
 * [49] Group Anagrams
 */

// @lc code=start

import java.util.ArrayList;
import java.util.List;
import java.util.Arrays;
import java.util.Map;
import java.util.HashMap;

class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> result = new ArrayList<>();
        for(String str: strs) {
            boolean anew = true;
            for(List<String> list: result) {
                if(isAnagram(str, list.get(0))) {
                    list.add(str);
                    anew = false;
                    break;
                }
            }
            
            if(anew) {
                List<String> list = new ArrayList<>();
                list.add(str);
                result.add(list);
            }
        }
        
        return result;
    }
    
    public boolean isAnagram(String s, String t) {
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

    public List<List<String>> groupAnagrams2(String[] strs) {
        char[][] chars = new char[strs.length][];
        for(int i = 0, len = strs.length; i < len; i++) {
            chars[i] = strs[i].toCharArray();
            Arrays.sort(chars[i]);
        }

        List<List<Integer>> indexGroup = new ArrayList<>();
        for(int k = 0, len = strs.length; k < len; k++) {
            boolean anew = true;
            for(List<Integer> list: indexGroup) {
                if(Arrays.equals(chars[list.get(0)], chars[k])) {
                    list.add(k);
                    anew = false;
                    break;
                }
            }

            if(anew) {
                List<Integer> list = new ArrayList<>();
                list.add(k);
                indexGroup.add(list);
            }
        }

        List<List<String>> result = new ArrayList<>();
        for(List<Integer> index: indexGroup) {
            List<String> vals = new ArrayList<>();
            for(Integer i: index) {
                vals.add(strs[i]);
            }
            result.add(vals);
        }


        return result;
    }

    public List<List<String>> groupAnagrams3(String[] strs) {
        
        Map<String, List<String>> map = new HashMap<>();
        List<List<String>> result = new ArrayList<>();
        for(String str: strs) {
            String anagram = getAnagram(str);
            if(map.containsKey(anagram)) {
                map.get(anagram).add(str);
            } else {
                List<String> list = new ArrayList<>();
                list.add(str);
                map.put(anagram, list);
                result.add(list);
            }
        }
        
        return result;
    }
    
    private String getAnagram(String str) {
        char[] chars = new char[26];
        for(char c: str.toCharArray()) {
            chars[c - 'a']++;
        }
        
        return new String(chars);
    }
}
// @lc code=end

