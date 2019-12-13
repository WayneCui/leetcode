/*
 * @lc app=leetcode id=859 lang=java
 *
 * [859] Buddy Strings
 */

// @lc code=start
import java.util.*;
class Solution {
    public boolean buddyStrings(String A, String B) {
        int n1 = A.length();
        int n2 = B.length();
        if(n1 != n2) { return false; }

        Set<Character> set = new HashSet<>();
        List<Integer> diffCharIndices = new ArrayList<>();
        for(int i = 0; i < n1; i++) {
            set.add(A.charAt(i));

            if(A.charAt(i) != B.charAt(i)) {
                diffCharIndices.add(i);
            }
        }

        if(diffCharIndices.size() == 0) {
            return set.size() != n1;
        } else if(diffCharIndices.size() == 2) {
            int first = diffCharIndices.get(0);
            int second = diffCharIndices.get(1);
            if(A.charAt(first) == B.charAt(second) && A.charAt(second) == B.charAt(first)) {
                return true;
            }
        }
        
        return false;
    }
}
// @lc code=end

