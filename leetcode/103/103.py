# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        layers = [[root], []]
        res = [[]]
        layer = 0
        
        while len(layers[0]) > 0:
            while len(layers[0]) > 0:
                t = layers[0][0]
                layers[0] = layers[0][1:]
                if t is None:
                    continue
                res[layer].append(t.val)
                if t.left is not None:
                    layers[1].append(t.left)
                if t.right is not None:            
                    layers[1].append(t.right)
            layers[0], layers[1] = layers[1], layers[0]
            layer += 1
            res.append([])

        res = [
            list(reversed(l)) if idx & 1 else l
            for idx, l in enumerate(res) 
            if len(l) > 0   
        ] 
        return res


