/*
 * @lc app=leetcode id=40 lang=java
 *
 * [40] Combination Sum II
 */

// @lc code=start
import java.util.*;

class Solution {
    private List<List<Integer>> result = new ArrayList<>();

    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<Integer> solution = new ArrayList<>();
        search(candidates, target, solution);

        return result;
    }

    private void search(int[] candidates, int target0, List<Integer> solution) {
        for(int i = 0, len = candidates.length; i < len; i++) {
            int c = candidates[i];
            if(i > 0 && c == candidates[i - 1]) { continue; }

            List<Integer> solution0 = new ArrayList<>(solution);
            if(target0 - c == 0) {
                solution0.add(c);
                this.result.add(solution0);
                return;
            } else if(target0 - c > 0) {
                solution0.add(c);
                search(Arrays.copyOfRange(candidates, i + 1, len), target0 - c, solution0);
            }
        }
    }
}
// @lc code=end

