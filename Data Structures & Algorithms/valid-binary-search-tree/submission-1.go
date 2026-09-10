/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}
func valid(node *TreeNode, left, right int)bool{
	if node==nil{
		return true
	}
	nodeval:=node.Val
	if nodeval<=left || nodeval>=right{
		return false
	}
	return valid(node.Left, left, nodeval) && valid(node.Right, nodeval, right)
}