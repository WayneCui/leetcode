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

    public List<List<Integer>> threeSum2(int[] nums) {
        int len = nums.length;
        List<List<Integer>> result = new ArrayList<>();

        if(len < 3) { return result; }

        Arrays.sort(nums);

        for(int i = 0; i < len - 2; i++) {
            if(i > 0 && nums[i] == nums[i - 1]) { continue; }
            int target = 0 - nums[i];
            int low = i + 1, high = len - 1;
            while(low < high) {
                if(nums[low] + nums[high] > target) {
                    do {
                        high--;
                    } while (high < len - 1 && high > i + 1 && nums[high] == nums[high + 1]);
                } else if(nums[low] + nums[high] < target) {
                    do {
                        low++;
                    } while (low > i + 1 && low < len - 1 && nums[low] == nums[low - 1]);
                } else {
                    result.add(Arrays.asList(nums[i], nums[low], nums[high]));

                    do {
                        high--;
                    } while (high < len - 1 && high > i + 1 && nums[high] == nums[high + 1]);

                    do {
                        low++;
                    } while (low > i + 1 && low < len - 1 && nums[low] == nums[low - 1]);
                }
            }
        }

        return result;
    }

    public List<List<Integer>> threeSum3(int[] nums) {
        int len = nums.length;
        List<List<Integer>> result = new ArrayList<>();

        if(len < 3) { return result; }

        Arrays.sort(nums);
        Map<Integer, Integer> next = new HashMap<>();
        next.put(nums[len - 1], len);
        Map<Integer, Integer> prev = new HashMap<>();
        prev.put(nums[0], -1);

        int ni = 0, pi = 0;
        for(int k = 0; k < len; k++) {
            if(k > 0 && nums[k] != nums[k - 1]) {
                prev.put(nums[k], pi);

                ni = k;
                next.put(nums[k - 1], ni);
            }
            pi = k;
        }

        for(int i = 0; i < len - 2; i++) {
            if(i > 0 && nums[i] == nums[i - 1]) { continue; }
            int target = 0 - nums[i];
            int low = i + 1, high = len - 1;
            while(low < high) {
                if(nums[low] + nums[high] > target) {
                    high = prev.get(nums[high]);
                } else if(nums[low] + nums[high] < target) {
                    low = next.get(nums[low]);
                } else {
                    result.add(Arrays.asList(nums[i], nums[low], nums[high]));

                    high = prev.get(nums[high]);
                    low = next.get(nums[low]);
                }
            }
        }

        return result;
    }
}
// @lc code=end

