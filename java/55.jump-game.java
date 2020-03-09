/*
 * @lc app=leetcode id=55 lang=java
 *
 * [55] Jump Game
 */

// @lc code=start
class Solution {
    public boolean canJump(int[] nums) {
        int len = nums.length;
        if(len == 0 || len == 1) {
            return true;
        }

        return canJump(nums, 0);

    }

    public boolean canJump(int[] nums, int current) {
        int len = nums.length;
        if(current == len - 1) {
            return true;
        }

        boolean result = false;
        for(int i = 1; i <= Math.min(nums[current], len - 1 - current); i++) {
            result = result || canJump(nums, current + i);
        }

        return result;
    }

    // greedy
    public boolean canJump2(int[] nums) {
        int reached = 0;
        for(int i = 0; i <= reached; i++) {
            reached = Math.max(reached, i + nums[i]);
            if(reached >= nums.length - 1) {
                return true;
            }
        }
        
        return false;
    }
}
// @lc code=end

