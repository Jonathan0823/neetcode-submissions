/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil { 
		return nil
	}

	dummy := &ListNode{}
	dummy.Next = head
	slow, fast := dummy, head

	for n > 0 { 
		fast = fast.Next
		n--
	}
    
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
