package lianbiao
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    dummy2 := &ListNode{0, headA}
    dummy1 := &ListNode{0, headB}
	
	cur1:= dummy2.Next
	cur2 := dummy1.Next
	//测量curl2和cur1长度
	i:=0
	i1:=0
	j:=1
	for cur2!=nil{
		cur2=cur2.Next
		i++
	}
	lengthcur2:=i
	for cur1!=nil{
		cur1=cur1.Next
		i1++
	}
	//反转cur1 	cur2
	
	rcur1:=reverseList(cur1)
	rcur2:=reverseList(cur2)
	for rcur1!=nil && rcur2!=nil{
	    rcur1=rcur1.Next
		rcur2=rcur2.Next
		
		if rcur1!=rcur2{
			break
		}
		j++
		
	}
	if j==i || j==i1{
	   return nil
	}
	
	    
    dummy := &ListNode{0, cur2}
    cur := dummy
	if lengthcur2-j<=0{
		return nil
	}
    for k:=0;k<lengthcur2-j;k++{
        cur = cur.Next
    }
    jiaodian:=cur.Next
    return jiaodian
	}
