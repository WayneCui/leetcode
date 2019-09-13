/*
 * @lc app=leetcode id=119 lang=rust
 *
 * [119] Pascal's Triangle II
 */
impl Solution {
    pub fn get_row(row_index: i32) -> Vec<i32> {
        let mut layer = vec![1];
        for num in (1..row_index + 1) {
            let mut nextLayer: Vec<i32> = Vec::new();
            
            let mut prev = &0;
            for i in &layer {
                nextLayer.push(prev + i);
                prev = i;
            }
            nextLayer.push(1);
            
            layer = nextLayer;
        }
        
        layer
    }
}

impl Solution {
    pub fn get_row(row_index: i32) -> Vec<i32> {
        let mut layer = vec![1; (row_index + 1) as usize];
        layer[0] = 1;
        for num in (1..row_index + 1) {            
            let mut prev = 0;
            let mut tmp = 0;
            for i in (0..num) {
                tmp = layer[i as usize];
                layer[i as usize] = layer[i as usize] + prev;
                prev = tmp
            }            
        }        
        layer
    }
}
