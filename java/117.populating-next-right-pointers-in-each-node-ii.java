/*
 * @lc app=leetcode id=117 lang=java
 *
 * [117] Populating Next Right Pointers in Each Node II
 */
/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}

    public Node(int _val,Node _left,Node _right,Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
*/
class Solution {
    public Node connect(Node root) {
        if(root == null) {
            return root;
        }
        
        List<Node> level = new ArrayList<>();
        level.add(root);
        while(level.size() > 0) {
            List<Node> nextLevel = new ArrayList<>();
            for(int i = 0; i < level.size() - 1; i++) {
                Node node = level.get(i);
                Node next = level.get(i + 1);
                node.next = next;

                if(node.left != null) {
                    nextLevel.add(node.left);
                }
                
                if(node.right != null) {
                    nextLevel.add(node.right);
                }
            }

            Node node = level.get(level.size() - 1);
            if(node.left != null) {
                nextLevel.add(node.left);
            }
            
            if(node.right != null) {
                nextLevel.add(node.right);
            }

            level = nextLevel;
        }

        return root;
    }
}

