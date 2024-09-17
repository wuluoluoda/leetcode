#include <unordered_map>
struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};
class Solution
{
public:
    ListNode *deleteDuplicates(ListNode *head)
    {
        // 遍历链表
        // 申请ya节点
        ListNode* preHead = new ListNode(0);
        preHead->next = head;
        
        ListNode *ya = preHead;
        // 申请map
        unordered_map<int, bool> m = {} while (ya->next != nullptr)
        {

            if (m[ya->next->val])
            {
                ya->next = ya->next->next;
            }
            else
            {

                m[ya->next->val] = true;
                ya = ya->next;
            }
        }
        // map记录
        // 若已存在，删除
        return preHead->next;
    }
};
