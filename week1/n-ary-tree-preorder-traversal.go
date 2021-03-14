/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
//前序遍历顺序为根左右
func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var ans []int
	ans = append(ans, root.Val)
	for _, node := range root.Children {
		ans = append(ans, preorder(node)...)
	}
	return ans
}
