/**12func reorderList(head *ListNode) {
$0
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// first point to last, second stay and point to last -1
//0 1 2 3 4 5 6 
// 0 6 1 5 2 4
// 0->1
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
        return
    }
	slow, fast:= head, head
	for fast!=nil && fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
	}
	mid:=slow.Next
	slow.Next=nil
	var prev *ListNode
	for mid!=nil{
		next:=mid.Next
		mid.Next=prev
		prev=mid
		mid=next
	}
	first:=head
	second:=prev
	for second!=nil {
		temp1, temp2:=first.Next, second.Next
		first.Next=second
		second.Next=temp1
		first=temp1
		second=temp2
	}
}
