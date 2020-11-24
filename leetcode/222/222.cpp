/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
struct TreeNode {
    int              val;
    struct TreeNode *left;
    struct TreeNode *right;
};

class Solution {
  public:
    int countNodes(TreeNode *root) {
        if (root == 0)
            return 0;
        return 1 + countNodes(root->left) + countNodes(root->right);
    }
};