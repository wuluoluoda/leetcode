#include <stack>
class Solution
{
public:
    bool isValid(string s){
        // 初始化一个map
        unordered_map<char, char> pairs = {
            {')', '('},
            {']', '['},
            {'}', '{'}
        };


        // 初始化一个栈*******
        stack<char> stk;
        // len奇数直接false
        if (s.size()% 2 == 1) {
                       return false;
        } // 用栈来处理好一些
                                    
 // 写一个map

        

        // 遍历
        for (int i= 0; i < s.size(); i++){

            // 如果得到右括号，就出栈
            // map存在代表着有左括号，出栈
            // map存在意味着map值大于0（如果没有值byte为空）
            if (pairs[s[i]] > 0){
                // 不能出栈或右括号与栈中左括号不对应********这里体现c的不同
                if (stk.empty() || stk.top() != pairs[s[i]])
                {
                    return false;
                }

                // 出栈
                stk.pop();
                }
            else{ // 如果得到了一个左括号，就入栈
               stk.push(s[i]);
            }
        }
        // 栈中无数即可
        return stk.size() == 0;
    }
};