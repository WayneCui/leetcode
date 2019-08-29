/*
 * @lc app=leetcode id=1 lang=rust
 *
 * [1] Two Sum
 */

use std::collections::HashMap;

impl Solution {
    pub fn two_sum0(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut dict = HashMap::new();

        let n = nums.len();
        for i in 1..n {
            dict.insert(nums[i], i);
        }

        let mut output = Vec::new();
        output.push(-1);
        output.push(-1);
        for i in 1..n {
            let first = nums[i];
            let second = target - first;
            let exist = dict.get(&second);

            match exist {
                Some(x) => {
                    if *x != i {
                        output[0] = i;
                        output[1] = *x;
                        break; 
                    }
                },
                None => {}
            }
        }

        return output;
    }

    // a more succinct version
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut dict = HashMap::new();
        
        for (i, num) in nums.iter().enumerate() {
            if let Some(&idx) = dict.get(&(target - num)) {
                return vec![idx, i as i32];
            } else {
                dict.insert(num, i as i32);
            }
        }

        vec![-1, -1]
    }
}

