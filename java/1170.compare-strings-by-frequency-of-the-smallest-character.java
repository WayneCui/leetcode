/*
 * @lc app=leetcode id=1170 lang=java
 *
 * [1170] Compare Strings by Frequency of the Smallest Character
 */

// @lc code=start
import java.util.*;
class Solution {
    public int[] numSmallerByFrequency(String[] queries, String[] words) {
        int[] queryFos = new int[queries.length];
        for(int i = 0; i < queries.length; i++) {
            queryFos[i] = fos(queries[i]);
        }

        int[] wordsFos = new int[words.length];
        for(int j = 0; j < words.length; j++) {
            wordsFos[j] = fos(words[j]);
        }
        Arrays.sort(wordsFos);
        System.out.println(Arrays.toString(wordsFos));
        System.out.println(Arrays.toString(queryFos));
        
        int[] result = new int[queries.length];
        for(int k = 0; k < queryFos.length; k++) {
            int index = binarySearch(wordsFos, queryFos[k] + 1);
            if(index < 0) {
                result[k] = wordsFos.length - (-index - 1);
            } else {
                result[k] = wordsFos.length - index;
            }
        }

        return result;
    }

    public int fos(String word) {
        int[] memo = new int[26];
        for(char c: word.toCharArray()) {
            memo[c - 'a'] += 1;
        }

        for(int count: memo) {
            if(count != 0) {
                return count;
            }
        }
        return 0;
    }

    int binarySearch(int[] a, int key) {
        int low = 0;
        int high = a.length;

        while (low < high) {
            int mid = (low + high) >>> 1;
            int midVal = a[mid];

            if (midVal < key)
                low = mid + 1;
            else
                high = mid;
        }

        if(low >= a.length || a[low] != key) {
            return -(low + 1);  // key not found.
        } else {
            return low;
        }
    }
}
// @lc code=end

