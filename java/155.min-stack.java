// credits to https://leetcode.com/problems/min-stack/discuss/49010/Clean-6ms-Java-solution
/*
 * @lc app=leetcode id=155 lang=java
 *
 * [155] Min Stack
 */
import java.util.*;
// @lc code=start
class MinStack {
    private LinkedList<Node> stack;

    /** initialize your data structure here. */
    public MinStack() {
        stack = new LinkedList<Node>();
    }
    
    public void push(int x) {
        Node node = new Node(x, Math.min(x, getMin()));
        stack.addFirst(node);
    }
    
    public void pop() {
        stack.poll();
    }
    
    public int top() {
        if(stack.peek() != null) {
            return stack.peek().val;
        }
        throw new NullPointerException();
    }
    
    public int getMin() {
        if(stack.peek() == null) {
            return Integer.MAX_VALUE;
        }
        return stack.peek().min;
    }

    class Node {
        int val;
        int min;

        public Node(int val, int min) {
            this.val = val;
            this.min = min;
        }
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */
// @lc code=end

class MinStack {
    Node head;

    /** initialize your data structure here. */
    public MinStack() {
        
    }
    
    public void push(int x) {
        if(head == null) {
            head = new Node(x, x);
        } else {
            head = new Node(x, Math.min(x, getMin()), head);
        }
    }
    
    public void pop() {
        head = head.next;
    }
    
    public int top() {
        return head.val;
    }
    
    public int getMin() {
        return head.min;
    }

    class Node {
        int val;
        int min;
        Node next;

        public Node(int val, int min) {
            this.val = val;
            this.min = min;
        }

        public Node(int val, int min, Node next) {
            this.val = val;
            this.min = min;
            this.next = next;
        }
    }
}