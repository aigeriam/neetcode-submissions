/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
        return nil
    }
	q:=[]*TreeNode{root}
	res:=[]int{}
	for len(q)>0{
		level:=len(q)
		for i:=0; i<level; i++{
			node:=q[0]
			q=q[1:]
			if i == level-1 {
                res = append(res, node.Val)
            }
			if node.Left!=nil{
				q=append(q, node.Left)
			}
			if node.Right!=nil{
				q=append(q, node.Right)
			}
		}
	}
	return res
}
/// 1 2 3 he has children so take that, pop 3 , 