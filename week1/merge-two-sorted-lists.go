//采用迭代法，将较小的值放入结果中,若链表中还剩下有元素则追加到结果的末尾
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newListNode := &ListNode{}
	pre := newListNode
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			newListNode.Next = l2
			l2 = l2.Next
		} else {
			newListNode.Next = l1
			l1 = l1.Next
		}
		newListNode = newListNode.Next
	}
	if l1 != nil {
		newListNode.Next = l1
	}
	if l2 != nil {
		newListNode.Next = l2
	}

	return pre.Next
}

//采用递归法，递归的边界是其中有一个链表中没有了元素
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}