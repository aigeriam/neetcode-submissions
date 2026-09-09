/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
   if root==nil{
	return 0
   }
   var dfs func (node *TreeNode, max int)int
   res:=0
   dfs=func (node *TreeNode, max int)int{
	if node==nil{
		return 0
	}
	if node.Val>=max{
		res++
	}
	max=maximum(node.Val, max)
	dfs(node.Left, max)
	dfs(node.Right, max)
	return res
   }
   res=dfs(root, root.Val)
   return res

}
func maximum(a, b int) int {
    if a > b {
        return a
    }
    return b
}
