# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findTarget(self, root: TreeNode, k: int) -> bool:
        return self.dfs(root, k, root)
    
    def dfs(self, root, k, t):
        if t == None:
            return False
        v = t.val
        return (
            True 
            if k - v != v and self.get(root, k - v) else 
            (
                self.dfs(root, k, t.left) or 
                self.dfs(root, k, t.right)
            )
        )

    def get(self, root, target):
        t = root
        while 1:
            if t == None:
                return False
            if t.val == target:
                return True
            elif target < t.val:
                t = t.left
            else:
                t = t.right
