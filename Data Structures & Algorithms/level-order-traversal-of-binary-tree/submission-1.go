/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	res:=[][]int{}
	if root==nil{
		return [][]int{}
	}
	res=append(res, []int{root.Val})
	level:=[]*TreeNode{root}
	for len(level)!=0{
		levelcount:=len(level)
		for i:=0; i<levelcount; i++{
			cur:=level[0]
			level=level[1:]
			if cur.Left!=nil{
				level=append(level, cur.Left)
			}
			if cur.Right!=nil{
				level=append(level, cur.Right)
			}
		}
		arr:=[]int{}
		for _, r:=range level{
			arr=append(arr, r.Val)
		}
		if len(arr)!=0{
			res=append(res, arr)
		}
	}
	return res
}
	
	



