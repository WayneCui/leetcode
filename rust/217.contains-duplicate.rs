/*
 * @lc app=leetcode id=217 lang=rust
 *
 * [217] Contains Duplicate
 */
use std::collections::HashMap;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut memo = HashMap::new();
        for v in nums {
            if memo.contains_key(&v) {
                return true;
            } else {
                memo.insert(v, 1);
            }
        }
        
        false
    }
}

