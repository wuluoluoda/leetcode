package lianbiao

type MyLinkedList struct {
	Val  int
	Next *ListNode
}

//MyLinkedList() 初始化 MyLinkedList 对象。
func Constructor() MyLinkedList {
	return MyLinkedList{

		Next: &ListNode{},
		Val:  0,
	}
}

//int get(int index) 获取链表中下标为 index 的节点的值。
//如果下标无效，则返回 -1 。
func (this *MyLinkedList) Get(index int) int {
	p := this.Next
	if index < 0 {
		return -1
	}
	for i := 0; i <= index; i++ {

		if this.Next == nil && i != index {

			return -1
		}
		if i == index {
			return this.Val
		}
		p = p.Next //i=0,this 下标0
	}
	return -1
}

//void //addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。
//在插入完成后，新节点会成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	p := this.Next
	newNode := &ListNode{

		Next: this.Next,
		Val:  0,
	}
	newNode.Val = val
	newNode.Next = p.Next
	this.Next = newNode
}

//将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	p := this.Next
	newNode := &ListNode{

		Next: this.Next,
		Val:  0,
	}
	for p.Next != nil {
		p= p.Next
	}
	p.Next = newNode
	newNode.Next = nil
	newNode.Val = val
}

/*将一个值为 val 的节点插入到链表中下标为 index 的节点之前。
如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。
如果 index 比长度更大，该节点将 不会插入 到链表中。*/
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	p := this.Next
	newNode := &ListNode{

		Next: this.Next,
		Val:  0,
	}
if index ==this.Val+1 {
	this.AddAtTail(val)
	return
}
if index <0 {
	return
}
	if index >this.Val {
    return
}else{//插入到index节点
for i:= 0; i < index; i++ {
    for i!= index-1 {
       p= p.Next
    }
}
p.Next = newNode
	newNode.Next = nil
	newNode.Val = val
}

}

//如果下标有效，则删除链表中下标为 index 的节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	p := this.Next
	
	if index <0 {
		return
	}
		if index >this.Val {
		return
	}else{
		for i:= 0; i < index; i++ {
			for i!= index-1 {
			   p= p.Next
			}
		}

	}
	p.Next = p.Next.Next
}
