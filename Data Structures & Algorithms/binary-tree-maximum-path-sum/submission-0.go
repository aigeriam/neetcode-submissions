/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    var dfs func (*TreeNode)int
	res:=[]int{root.Val}
	dfs=func(node *TreeNode)int{
		if node==nil{
			return 0
		}
		leftmax:=dfs(node.Left)
		rightmax:=dfs(node.Right)
		leftmax=max(leftmax, 0)
		rightmax=max(rightmax, 0)
		res[0]=max(res[0], node.Val+leftmax+rightmax)
		return node.Val+max(leftmax, rightmax)
		
	}
	dfs(root)
	return res[0]
	}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}