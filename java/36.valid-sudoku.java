/*
 * @lc app=leetcode id=36 lang=java
 *
 * [36] Valid Sudoku
 */

// @lc code=start
class Solution {
    public boolean isValidSudoku(char[][] board) {
        for(int i = 0; i < 9; i++) {
            int[] row = new int[10];
            int[] col = new int[10];
            int[] block = new int[10];
            for(int j = 0; j < 9; j++) {
                if(board[i][j] != '.' && row[board[i][j] - '0']++ > 0) {
                    return false;
                }

                if(board[j][i] != '.' && col[board[j][i] - '0']++ > 0) {
                    return false;
                }

                int x = (i / 3) * 3 + (j / 3);
                int y = (i % 3) * 3 + (j % 3);
                if(board[x][y] != '.' && block[board[x][y] - '0']++ > 0) {
                    return false;
                }
            }
        }

        return true;
    }
}
// @lc code=end

