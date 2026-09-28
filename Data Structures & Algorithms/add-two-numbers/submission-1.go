/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 { 
		first, second := 0, 0
		if l1 != nil {
			first = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			second = l2.Val
			l2 = l2.Next
		}

		res := first + second + carry 
			cur.Next = &ListNode{Val: res%10}
			cur = cur.Next
			carry = res/10
	}

	return dummy.Next
}
