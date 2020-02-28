/*
 * @lc app=leetcode id=15 lang=java
 *
 * [15] 3Sum
 */

// @lc code=start

import java.util.*;
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        int len = nums.length;

        Arrays.sort(nums);
        if(len < 3 || nums[0] > 0 || nums[len - 1] < 0) {
            return result;
        }

        Map<Integer, Integer> cache = new HashMap<>(len);
        for(int i = 0; i < len; i++) {
            cache.put(nums[i], i);
        }

        for(int i = 0; i < len - 2; i++) {
            if(nums[i] > 0) {
                break;
            }

            if(i > 0 && nums[i - 1] == nums[i]) {
                continue;
            }

            int target = 0 - nums[i];
            for(int j = i + 1; j < len - 1; j++) {
                if(nums[j] > target) {
                    break;
                }

                if(j > i + 1 && nums[j - 1] == nums[j]) {
                    continue;
                }

                int target2 = target - nums[j];
                if(cache.containsKey(target2) && cache.get(target2) > j) {
                    result.add(Arrays.asList(nums[i], nums[j], target2));
                }
            }
        }

        return result;
    }
}
// @lc code=end

