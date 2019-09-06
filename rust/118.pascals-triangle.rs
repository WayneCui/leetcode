/*
 * @lc app=leetcode id=118 lang=rust
 *
 * [118] Pascal's Triangle
 */
impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        let mut res: Vec<Vec<i32>> = Vec::new();
        
        let mut layer = vec![1];
        for num in (1..num_rows + 1) {
            let mut nextLayer: Vec<i32> = Vec::new();
            
            let mut prev = &0;
            for i in &layer {
                nextLayer.push(prev + i);
                prev = i;
            }
            nextLayer.push(1);
            
            res.push(layer);
            layer = nextLayer;
        }
        
        res
    }
}

