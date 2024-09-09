package lianbiao

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyhead := &ListNode{Next: head}
	for p := dummyhead; p.Next != nil; {
		if p.Next.Val == val {//这是对字段的改变，比如p1的next变成了next。next
			p.Next = p.Next.Next
		} else { //这是对整体的改变，比如p1变成了p2，但是p1的next还是指向next
			//否则移动到下一个节点
			/*似乎这里.Next = p.Next.Next：
						这个操作是在当前节点 p 指向的下一个节点需要被删除时执行的。
						如果 p 指针是哑节点 dummyhead，那么这个操作会改变哑节点的 Next 指针，
						使其跳过被删除的头节点，指向下一个节点。如果 p 不是哑节点，这个操作会改变 p 所在位置的链表结构，
						但不影响哑节点的 Next 指针。

			p = p.Next：这个操作是在遍历链表时，将 p 指针移动到下一个节点。
			这个操作不会改变任何节点的 Next 指针，它只是改变了 p 指针的值，
			使其指向下一个节点。这个操作对哑节点 dummyhead 的 Next 指针没有任何影响。*/
			p = p.Next
		}
	}
	return dummyhead.Next
}
