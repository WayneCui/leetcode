/*
 * @lc app=leetcode id=39 lang=java
 *
 * [39] Combination Sum
 */

// @lc code=start
import java.util.*;

class Solution {
    private List<List<Integer>> result = new ArrayList<>();
    
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<Integer> solution = new ArrayList<>();
        search(candidates, target, solution);

        return result;
    }

    private void search(int[] candidates, int target0, List<Integer> solution) {
        for(int i = 0, len = candidates.length; i < len; i++) {
            int c = candidates[i];
            List<Integer> solution0 = new ArrayList<>(solution);
            if(target0 - c == 0) {
                solution0.add(c);
                this.result.add(solution0);
                return;
            } else if(target0 - c > 0) {
                solution0.add(c);
                search(Arrays.copyOfRange(candidates, i, len), target0 - c, solution0);
            }
        }
    }
}
// @lc code=end

