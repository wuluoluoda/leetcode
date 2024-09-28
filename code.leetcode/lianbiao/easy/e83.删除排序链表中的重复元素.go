package lianbiao

func deleteDuplicates(head *ListNode) *ListNode {
	//遍历链表
	//申请ya节点
	dummy := &ListNode{Val: 0, Next: head}
	ya := dummy
	//申请map
	m := make(map[int]bool)
	for ya.Next != nil {

		if m[ya.Next.Val] {
			ya.Next = ya.Next.Next
		} else {

			m[ya.Next.Val] = true
			ya = ya.Next
		}

	}
	//map记录
	//若已存在，删除
	return dummy.Next
}
