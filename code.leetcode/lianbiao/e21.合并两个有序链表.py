 
class ListNode(object):
    def __init__(self, val=0, next=None):
       self.val = val
       self.next = next
class Solution(object):
	def mergeTwoLists(self, l1: ListNode, l2: ListNode):
		preHead = ListNode(-1) # 创建一个虚拟头节点
		prev = preHead               # 使用 prev 指针追踪当前应该连接的位置

		while l1 and l2 :
			if l1.Val <= l2.Val :
				prev.next = l1
				l1 = l1.next
			else: 
				prev.next = l2
				l2 = l2.next
		
			prev = prev.next # 移动 prev 指针

		if l1 :
			prev.next = l1
		else: 
			prev.next = l2

		return preHead.next
