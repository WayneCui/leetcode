/*
 * @lc app=leetcode id=18 lang=java
 *
 * [18] 4Sum
 */

// @lc code=start
import java.util.*;
class Solution {
    public List<List<Integer>> fourSum(int[] nums, int target) {
        int len = nums.length;
        List<List<Integer>> result = new ArrayList<>();

        if(len < 4) { return result; }

        Arrays.sort(nums);
        Map<Integer, Integer> map = new HashMap<>();
        for(int a = 0; a < len; a++) {
            map.put(nums[a], a);
        }

        for(int i = 0; i < len - 3; i++) {
            if(i > 0 && nums[i] == nums[i - 1]) { continue; }

            for(int j = i + 1; j < len - 2; j++) {
                if(j > i + 1 && nums[j] == nums[j - 1]) { continue; }

                int target0 = target - nums[i] - nums[j];
                for(int k = j + 1; k < len - 1; k++) {
                    if(k > j + 1 && nums[k] == nums[k - 1]) { continue; }

                    if(map.containsKey(target0 - nums[k]) &&
                            map.get(target0 - nums[k]) > k) {
                        result.add(Arrays.asList(nums[i], nums[j], nums[k], target0 - nums[k]));
                    }
                }
            }
        }

        return result;
    }
}
// @lc code=end

