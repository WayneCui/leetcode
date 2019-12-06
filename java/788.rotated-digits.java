/*
 * @lc app=leetcode id=788 lang=java
 *
 * [788] Rotated Digits
 */

// @lc code=start
class Solution {
    public int rotatedDigits(int N) {
        int count = 0;
        for (int i = 1; i <= N; i++) {
            if(isGood(i)) {
                count++;
            }
        }

        return count;
    }

    private boolean isGood(int n) {
        int rotated = rotate(n);
        if(rotated != -1 && rotated != n) {
            return true;
        }  

        return false;
    }

    private int rotate(int n) {
        String numStr = String.valueOf(n);
        char[] chars = numStr.toCharArray();
        char[] out = new char[numStr.length()];
        boolean validFlag = true;
        loop: for(int i = 0; i < numStr.length(); i++) {
            char c = chars[i];
            switch(c) {
                case '0':
                case '1':
                case '8':
                    out[i] = c;
                    break;
                case '2':
                    out[i] = '5';
                    break;
                case '5':
                    out[i] = '2';
                    break;
                case '6':
                    out[i] = '9';
                    break;
                case '9':
                    out[i] = '6';
                    break;
                default:
                    validFlag = false;
                    break loop;
                    // do nothing
            }
        }

        if(!validFlag) {
            return -1;
        } else {
            return Integer.parseInt(new String(out));
        }
    }
}
// @lc code=end

