/*
 * @lc app=leetcode id=796 lang=java
 *
 * [796] Rotate String
 */

// @lc code=start
class Solution {
    public boolean rotateString(String A, String B) {
        if(A.equals(B)) {
            return true;
        }
        int n1 = A.length();
        int n2 = B.length();
        if(n1 != n2) { return false; }

        int hashB = B.hashCode();
        for(int i = 0; i < n1; i++) {
            if(A.charAt(i) == B.charAt(0)) {
                int hashA = (A.substring(i) + A.substring(0, i)).hashCode();
                if(hashA == hashB) {
                    return true;
                }
            }
        }

        return false;
    }
}

class Solution {
    public boolean rotateString(String A, String B) {
        int n1 = A.length();
        int n2 = B.length();
        if(n1 != n2) { return false; }
        if(n1 == 0) { return true; }

        for(int i = 0; i < n1; i++) {
            if(A.charAt(i) == B.charAt(0)) {
                if(B.equals(A.substring(i) + A.substring(0, i))) {
                    return true;
                }
            }
        }

        return false;
    }
}

// @lc code=end

