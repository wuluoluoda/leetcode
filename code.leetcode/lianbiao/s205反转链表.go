package lianbiao
 type ListNode struct {
	     Val int
	    Next *ListNode
	 }
// 反转链表
   func reverseList(head *ListNode) *ListNode {
   var curl *ListNode
   var pre *ListNode
   curl=head
   for curl!=nil{
       next:=curl.Next
	   curl.Next=pre
	   pre=curl
	   curl=next
   }
 
  
return pre

   }