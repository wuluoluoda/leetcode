package lianbiao

//链表将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{Val: -1} // 创建一个虚拟头节点
	prev := prehead               // 使用 prev 指针追踪当前应该连接的位置

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next // 移动 prev 指针
	}
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return prehead.Next
}
