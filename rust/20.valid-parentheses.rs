/*
 * @lc app=leetcode id=20 lang=rust
 *
 * [20] Valid Parentheses
 */

use std::collections::HashMap;

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut dict:HashMap<&char, &char> = HashMap::new();
        dict.insert(&']', &'[');
        dict.insert(&'}', &'{');
        dict.insert(&')', &'(');

        let mut vec = Vec::new();
        for c in s.chars() {
            match c {
                '(' | '{' | '[' => { vec.push(c); },
                _ => {
                    let a = vec.pop();
                    let v = dict.get(&c);
                        
                    match a {
                        Some(x) => {
                            if x != **v.unwrap() {
                                return false;
                            }
                        },
                        None => return false,
                    }
                },
            }
        }

        vec.len() == 0
    }
}

