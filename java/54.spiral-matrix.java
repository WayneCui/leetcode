/*
 * @lc app=leetcode id=54 lang=java
 *
 * [54] Spiral Matrix
 */

// @lc code=start
import java.util.*;

class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        int row = matrix.length;
       
        if(row == 0) {
            return Collections.emptyList();
        }
        
        int col = matrix[0].length;
        if(col == 0) {
            return Collections.emptyList();
        }

        List<Integer> result = new ArrayList<>();
        int times = Math.min(row, col) / 2;
        for(int i = 0; i < times; i++) {
            for(int j = i; j < col - 1 - i; j++) {
                result.add(matrix[i][j]);
            }

            for(int k = i; k < row - 1 - i; k++) {
                result.add(matrix[k][col - 1 - i]);
            }

            for(int p = col - 1 - i; p > i; p--) {
                result.add(matrix[row - 1 - i][p]);
            }

            for(int q = row - 1 - i; q > i; q--) {
                result.add(matrix[q][i]);
            }
        }

        if(Math.min(row, col) % 2 == 1) {
            if(row > col) {
                for(int i = times; i < row - times; i++) {
                    result.add(matrix[i][times]);
                }
            } else if(row < col) {
                for(int i = times; i < col - times; i++) {
                    result.add(matrix[times][i]);
                }
            } else {
                result.add(matrix[times][times]);
            }
        }

        return result;
    }
}
// @lc code=end

