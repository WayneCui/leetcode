import java.util.ArrayList;

/*
 * @lc app=leetcode id=297 lang=java
 *
 * [297] Serialize and Deserialize Binary Tree
 */
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        if(root == null) {
            return "";
        }
        
        StringBuilder buf = new StringBuilder();
        List<TreeNode> level = new ArrayList<>();
        level.add(root);
        buf.append(root.val + ",");
        while(level.size() > 0) {
            List<TreeNode> nextLevel = new ArrayList();
            for(TreeNode node: level) {
                if(node.left != null) {
                    buf.append(node.left.val + ",");
                    nextLevel.add(node.left);
                } else {
                    buf.append(",");
                }

                if(node.right != null) {
                    buf.append(node.right.val + ",");
                    nextLevel.add(node.right);
                } else {
                    buf.append(",");
                }
            }

            level = nextLevel;
        }

        return buf.toString();
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        String[] vals = data.split(",");
        //System.out.println(Arrays.asList(vals));
        if("".equals(vals[0])) {
            return null;
        }

        TreeNode root = new TreeNode(Integer.parseInt(vals[0]));
        List<TreeNode> level = new ArrayList<>();
        level.add(root);

        int index = 1;
        while(index < vals.length && level.size() > 0){
            List<TreeNode> nextLevel = new ArrayList<>();
            for(TreeNode node: level) {
                if(index < vals.length && !"".equals(vals[index])){
                    TreeNode ln = new TreeNode(Integer.parseInt(vals[index]));
                    node.left = ln;
                    nextLevel.add(ln);
                }

                if(index + 1 < vals.length && !"".equals(vals[index + 1])){
                    TreeNode rn = new TreeNode(Integer.parseInt(vals[index + 1]));
                    node.right = rn;
                    nextLevel.add(rn);
                }
                
                index += 2;
            }
            
            level = nextLevel;
        }

        return root;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));

