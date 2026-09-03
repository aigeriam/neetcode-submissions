/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head:=&ListNode{} 
	merged := head
	for list1!=nil && list2!=nil{
		if list1.Val<list2.Val{
			merged.Next=list1
			list1=list1.Next
		}else{
			merged.Next=list2
			list2=list2.Next
		}
		merged=merged.Next 
	}
	merged.Next=list1
	if list1==nil{
		merged.Next=list2
	}
	return head.Next
}
