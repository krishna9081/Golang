/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return dfsHelper(root, 0)

}

func dfsHelper(node *TreeNode, sum int) int {
	//Leaf node
	if node.Left == nil && node.Right == nil {
		return node.Val + sum
	} else if node.Left == nil {
		sum += node.Val
		return dfsHelper(node.Right, sum*10)
	} else if node.Right == nil {
		sum += node.Val
		return dfsHelper(node.Left, sum*10)
	} else {
		sum += node.Val
		return dfsHelper(node.Left, sum*10) + dfsHelper(node.Right, sum*10)
	}

} 