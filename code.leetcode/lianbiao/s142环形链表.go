package lianbiao
type ListNode struct {
	Val  int
	Next *ListNode
}

//142. 环形链表 II


    //哈希表示例
	/*func reverseList(head *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
    for tmp := headA; tmp != nil; tmp = tmp.Next {
        vis[tmp] = true
    }
    for tmp := headB; tmp != nil; tmp = tmp.Next {
        if vis[tmp] {
            return tmp
        }
    }
    return nil
}

	*/
	//遍历链表
	func detectCycle(head *ListNode) *ListNode {
		//大框架哈希表检验
		//1.确定索引
		//2.返回tmp
		//先判断是否有环
		vis := map[*ListNode]bool{}
		for tmp := head; tmp != nil; tmp = tmp.Next {
			
			if vis[tmp] {
				return tmp
			}
			vis[tmp] = true
		}	
		return nil
}
