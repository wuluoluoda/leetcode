class Solution(object):
    class ListNode:
        def __init__(self, val=0, next=None):
            self.val = val
            self.next = next
    def deleteDuplicates(self, head):
        #遍历链表
        #申请ya节点
        dummy = self.ListNode(0,head)
        ya = dummy
        #申请map
        m = {}
        while ya.next is not None :
            if ya.next.val in m:
                ya.next = ya.next.next
            else :

                m[ya.next.val] = True
                ya = ya.next
            

        
        #map记录
        #若已存在，删除
        return dummy.next