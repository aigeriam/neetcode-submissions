/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tail:=head
    lenght:=0
    for tail!=nil{
        tail=tail.Next
        lenght++
    }
    goal:=lenght-n
    start:=head
    count:=0
    if goal==0{
        return head.Next
    }
    for count!=lenght{
        count++
        if count==goal{
            skipnode:=start.Next.Next
            start.Next=skipnode
            return head
        }
        start=start.Next
    }
    return head
}
