package lianbiao

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

图示两个链表在节点 c1 开始相交：
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//建立map
	m := make(map[*ListNode]bool)
	//遍历headA，将节点放入map
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}
	//遍历headB，如果节点在map中，则返回
	for headB != nil {
		if m[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
