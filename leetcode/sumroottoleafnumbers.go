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
	//Leaf node - All that leaf node is doing is adding Value to the actual sum
	if node.Left == nil && node.Right == nil {
		return node.Val + sum
	} else if node.Left == nil { // recursive cases value is added to sum and the sum is multiplied by 10
		sum += node.Val
		return dfsHelper(node.Right, sum*10)
	} else if node.Right == nil {
		sum += node.Val
		return dfsHelper(node.Left, sum*10)
	} else { //special recursive case when both the children exist ...sums of both the children get added
		sum += node.Val
		return dfsHelper(node.Left, sum*10) + dfsHelper(node.Right, sum*10)
	}

} 