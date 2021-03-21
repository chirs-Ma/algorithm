//首先根据前序遍历确定根节点，然后在中序遍历中找到根节点位置，最后递归获取左右子数
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{preorder[0],nil,nil}
    num := 0
    for k,v := range(inorder) {
        if v == preorder[0] {
            num = k
            break
        }
    }
    root.Left = buildTree(preorder[1:num+1],inorder[:num])
    root.Right = buildTree(preorder[num+1:],inorder[num+1:])
    return root
}