/*
 * @lc app=leetcode id=520 lang=rust
 *
 * [520] Detect Capital
 */

// @lc code=start
impl Solution {
    pub fn detect_capital_use(word: String) -> bool {
        let bytes = word.as_bytes();
        let mut state = -1;
        for(i, &a_byte) in bytes.iter().enumerate() {
            match state {
                0 | 1 => {
                    if a_byte <= b'Z' {
                        return false;
                    }
                },
                2 => {
                    if a_byte >= b'a' {
                        return false;
                    }
                },
                _ => {
                    if i == 0 {
                        if a_byte >= b'a' && a_byte <= b'z' {
                            state = 0;
                        }
                    } else {
                        if a_byte >= b'a' && a_byte <= b'z' {
                            state = 1;
                        } else {
                            state = 2;
                        }
                    }
                }
            }
        }

        true
    }
}
// @lc code=end

