

 struct ListNode{
      int val;
      ListNode *next;
      ListNode() : val(0), next(nullptr) {}
      ListNode(int x) : val(x), next(nullptr) {}
      ListNode(int x, ListNode *next) : val(x), next(next) {}
 };

class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
    ListNode* preHead =new ListNode(-1); // 创建一个虚拟头节点
	
    ListNode* prev = preHead;               // 使用 prev 指针追踪当前应该连接的位置

		while(l1!=nullptr && l2!=nullptr) {
			if (l1->val <= l2->val) {
				prev->next = l1;
				l1 = l1->next;}
			else{ 
				prev->next = l2;
				l2 = l2->next;}
		
			prev = prev->next;
			} // 移动 prev 指针

		if (l1!=nullptr) {
			prev->next = l1;}
		else{ 
			prev->next = l2;}

		return preHead->next ;   
    }
};