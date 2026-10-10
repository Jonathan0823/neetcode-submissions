/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	slow, fast := head, head.Next

	for fast != nil { 
		slow = slow.Next
		fast = fast.Next
		if fast != nil { 
			fast = fast.Next
		}
	}

	second := slow.Next
	slow.Next = nil 
	cur := second
	var prev *ListNode
	for cur != nil { 
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	first := head
	second = prev
	for second != nil { 
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}
