/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */
 func reverseVowels(s string) string {
    bytes := []byte(s)
    n := len(bytes)
    i, j := 0, n - 1
    for i < j {
        for i < n {
            switch bytes[i] {
            case 'a','e','i','o','u','A', 'E', 'I', 'O', 'U':
                break
            }
            i++
        }
        
        for j >= 0 {
            switch bytes[J] {
            case 'a','e','i','o','u','A', 'E', 'I', 'O', 'U':
                break
            }
            j--
        }
        
        if i < j {
            bytes[i], bytes[j] = bytes[j], bytes[i]
            i++
            j--
        }
    }
    
    return string(bytes)
}

