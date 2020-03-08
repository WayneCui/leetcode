/*
 * @lc app=leetcode id=47 lang=java
 *
 * [47] Permutations II
 */

// @lc code=start
import java.util.*;

class Solution {
    private List<List<Integer>> result = new ArrayList<>();
    private int[] nums;
    private int len;

    public List<List<Integer>> permuteUnique(int[] nums) {
        Arrays.sort(nums);
        this.nums = nums;
        this.len = nums.length;

        List<Integer> p = new ArrayList<>();
        boolean[] used = new boolean[len];
        backtrack(used, p);

        return result;
    }

    private void backtrack(boolean[] used, List<Integer> p) {
        if(p.size() == this.len) {
            this.result.add(new ArrayList<>(p));
            return;
        }

        for(int i = 0; i < this.len; i++) {

            if(used[i]) {
                continue;
            }

            if((i > 0 && this.nums[i] == this.nums[i - 1]) && !used[i - 1]) {
                continue;
            }

            p.add(this.nums[i]);
            used[i] = true;
            backtrack(used, p);
            used[i] = false;
            p.remove(p.size() - 1);
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] nums = {1,1,2};
        System.out.println(solution.permuteUnique(nums));
    }
}
// @lc code=end

