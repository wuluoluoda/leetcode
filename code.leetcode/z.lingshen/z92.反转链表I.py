class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        p0 = dummy = ListNode(next=head)
        for _ in range(left - 1):
            p0 = p0.next

        pre = None
        cur = p0.next
        for _ in range(right - left + 1):
            nxt = cur.next
            cur.next = pre  # 每次循环只修改一个 next，方便大家理解
            pre = cur
            cur = nxt

        # 见视频
        p0.next.next = cur
        p0.next = pre
                    ``````````````                                                                  ``                                                                  ··       
        return dummy.next
'''p0.next.next = cur：将反转子链表的最后一个节点（p0.next）的 next 指针指向 cur，
即原链表中反转子链表后面的部分。
p0.next = pre：将 p0 的 next 指针指向 pre，
即反转子链表的第一个节点，这样整个链表就连接起来了。'''   

