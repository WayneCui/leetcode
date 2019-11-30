/*
 * @lc app=leetcode id=929 lang=java
 *
 * [929] Unique Email Addresses
 */

// @lc code=start
import java.util.*;
class Solution {
    public int numUniqueEmails(String[] emails) {
        Set<String> set = new HashSet<String>();
        for(String email: emails) {
            String[] components = email.split("@");
            String localname = components[0];
            String domain = components[1];
            String[] parts = localname.replaceAll("\\.", "").split("\\+");
            String resolved = parts[0] + "@" + domain;
            set.add(resolved);
        }

        return set.size();
    }
}
// @lc code=end


class Solution {
    public int numUniqueEmails(String[] emails) {
        Set<String> set = new HashSet<String>();
        for(String email: emails) {
            String resolved = resolve(email);
            set.add(resolved);
        }

        return set.size();
    }

    public String resolve(String email) {
        char[] chars = email.toCharArray();
        int i = 0, j = 0;
        boolean domain = false, ignore = false;
        while(j < chars.length) {
            switch(chars[j]) {
                case '.':
                    if(domain) {
                        chars[i++] = chars[j];
                    }
                    break;
                case '+':
                    if(domain) {
                        chars[i++] = chars[j];
                    } else {
                        ignore = true;
                    }
                    break;
                case '@':
                    domain = true;
                    chars[i++] = chars[j];
                    break;
                default:
                    if(domain || !ignore) {
                        chars[i++] = chars[j];
                    }
            }

            j++;
        }

        return new String(chars, 0, i);
    }
}
