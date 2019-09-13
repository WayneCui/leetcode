/*
 * @lc app=leetcode id=11 lang=rust
 *
 * [11] Container With Most Water
 */
impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut i: usize = 0;
        let mut j: usize = (height.len() - 1) as usize;

        let mut water = 0;
        while i < j {
            water = water.max((j - i) * height[i].min(height[j]) as usize);
            if height[i] < height[j] {
                i = i + 1;
            } else {
                j = j - 1;
            }
        }
        water as i32
    }
}

