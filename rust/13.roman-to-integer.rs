impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let mut result = 0;
        let n = s.len();
        let bytes = s.as_bytes();
        for (i, &a_byte) in bytes.iter().enumerate() {
            match a_byte {
                b'I' => { 
                    if i < n - 1 && (bytes[i+1] == b'V' || bytes[i+1] == b'X') {
                        result -= 1;
                    } else {
                        result += 1;
                    }
                },
                b'V' => { result += 5; },
                b'X' => {
                    if i < n - 1 && (bytes[i+1] == b'L' || bytes[i+1] == b'C') {
                        result -= 10;
                    } else {
                        result += 10;
                    }
                },
                b'L' => { result += 50; },
                b'C' => {
                    if i < n - 1 && (bytes[i+1] == b'D' || bytes[i+1] == b'M') {
                        result -= 100;
                    } else {
                        result += 100;
                    }
                },
                b'D' => { result += 500; },
                b'M' => { result += 1000; },
                _ => {}
            }
        }
        
        return result
    }
}