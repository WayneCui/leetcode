/*
 * @lc app=leetcode id=59 lang=java
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
class Solution {
    public int[][] generateMatrix(int n) {
        int[][] result = new int[n][n];
        int layer = n / 2;
        int num = 0;
        for(int L = 0; L < layer; L++) {
            for(int i = L; i < n - L - 1; i++) {
                result[L][i] = ++num;
            }
            
            for(int j = L; j < n - L - 1; j++) {
                result[j][n - L - 1] = ++num;
            }
            
            for(int p = n - L - 1; p > L; p--) {
                result[n - L - 1][p] = ++num;
            }
            
            for(int q = n - L - 1; q > L; q--) {
                result[q][L] = ++num;
            }
        }
        
        if(n % 2 == 1) {
            result[layer][layer] = ++num;
        }
        
        return result;
    }
}
// @lc code=end

