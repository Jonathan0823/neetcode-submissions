/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	l, r := dummy, head
	for n > 0 && r != nil{ 
		r = r.Next
		n--
	}

	for r != nil { 
		l = l.Next
		r = r.Next
	}

	l.Next = l.Next.Next

	return dummy.Next
    
}
