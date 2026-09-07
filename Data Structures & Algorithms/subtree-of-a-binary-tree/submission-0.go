/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root==nil{
		return false 
	}
	if isthesame(root, subRoot){
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isthesame(root *TreeNode, subRoot *TreeNode) bool {
	if root==nil && subRoot==nil{
		return true 
	}
	if root!=nil && subRoot!=nil && root.Val==subRoot.Val{
		return isthesame(root.Left, subRoot.Left) && isthesame(root.Right, subRoot.Right)
	}
	return false 
}
