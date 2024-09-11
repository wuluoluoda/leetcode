#include <unordered_map>
class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
    //建立map
  unordered_map<ListNode *, bool> m = {};
  //遍历headA，将节点放入map
  while (headA!=nullptr){
    m[headA] =true;  
    headA = headA->next;
  }
  //遍历headB，如果节点在map中，则返回
	while (headB!=nullptr){
		//map中存在headb的语法
        /*如果map或unordered_map中不存在headB这个键，
        find方法返回的迭代器等于end()迭代器，因此在使用返回的迭代器之前，
        应该检查它是否等于end()。*/
        if (m.find(headB) != m.end()){
			return headB;
		}
		headB = headB->next;
	}
  return nullptr ;   
    }
};