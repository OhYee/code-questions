/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
struct TreeNode {
    int              val;
    struct TreeNode *left;
    struct TreeNode *right;
};

int countNodes(struct TreeNode *root) {
    if (root == 0)
        return 0;
    return 1 + (root->left != 0 ? countNodes(root->left) : 0) +
           (root->right != 0 ? countNodes(root->right) : 0);
}