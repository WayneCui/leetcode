/* -----------------------------------
 *  WARNING:
 * -----------------------------------
 *  Your code may fail to compile
 *  because it contains public class
 *  declarations.
 *  To fix this, please remove the
 *  "public" keyword from your class
 *  declarations.
 */


/*
 * @lc app=leetcode id=148 lang=java
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode sortList(ListNode head) {
        if(head == null || head.next == null) return head;
        ListNode pivot = head;
        ListNode leftHead = new ListNode(-1);
        ListNode rightHead = new ListNode(-1);

        ListNode node = head.next;
        ListNode L = leftHead, R = rightHead, P = pivot;
        while(node != null) {
            if(node.val < pivot.val) {
                L.next = node;
                L = L.next;
            } else if(node.val > pivot.val){
                R.next = node;
                R = R.next;
            } else {
                P.next = node;
                P = P.next;
            }
            node = node.next;
        }
        L.next = null;
        R.next = null;
        
        ListNode sortedLeft = sortList(leftHead.next);
        ListNode sortedRight = sortList(rightHead.next);

        ListNode dummy = new ListNode(-1);
        if(sortedLeft == null) {
            dummy.next = pivot;
        } else {
            dummy.next = sortedLeft;
            ListNode e = sortedLeft;
            while (e != null && e.next != null) {
                e = e.next;
            }
            e.next = pivot;
        }
        
        P.next = sortedRight;

        return dummy.next;
    }
}


// @lc code=end

