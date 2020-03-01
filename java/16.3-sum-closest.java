/*
 * @lc app=leetcode id=16 lang=java
 *
 * [16] 3Sum Closest
 */

// @lc code=start
class Solution {
    public int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);
        int len = nums.length;
        if(len < 3) {
            int res = 0;
            for(int num: nums) {
                res += num;
            }
            return res;
        }

        int res = nums[0] + nums[1] + nums[2];
        for(int j = 0; j < len - 2; j++) {
            int low = j + 1;
            int high = len - 1;
            while(low < high) {
                int sum = nums[j] + nums[low] + nums[high];
                if(Math.abs(sum - target) <= Math.abs(res - target)) {
                    res = sum;
                    if(res == target) {
                        return res;
                    }
                }

                if(sum > target) {
                    high--;
                } else {
                    low++;
                }
            }
        }

        return res;
    }
}
// @lc code=end

