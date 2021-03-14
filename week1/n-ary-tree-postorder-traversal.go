/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
//后序遍历顺序为左右根
func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	results := []int{}
	for _, v := range root.Children {
		results = append(results, postorder(v)...)
	}

	results = append(results, root.Val)

	return results

}