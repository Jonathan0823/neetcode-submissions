/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	 dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil { 
		if list1.Val < list2.Val { 
			cur.Next = list1
			list1 = list1.Next
		} else { 
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	rest := list1
	if list2 != nil { 
		rest = list2
	}

	cur.Next = rest
	
	return dummy.Next
    
}
