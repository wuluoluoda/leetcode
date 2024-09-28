class Solution(object):
    def isValid(self, s):
       #zhan的初始化很简单
        stack=list()
        #len奇数直接False
        if len(s)%2==1:
            return False

        #用栈来处理好一些
        #写一个map
        pairs = {
        ')': '(',
        ']': '[',
        '{': '}',
        }

      

        #遍历
        i=0
        while i<len(s):
            

        #如果得到右括号，就出栈
            #map存在代表着有左括号，出栈
                #map存在意味着map值大于0（如果没有值byte为空）
            if s[i] in pairs:
                #不能出栈或右括号与栈中左括号不对应
                if not stack or stack.pop() != pairs[s[i]]:  
                    return False

                
                #出栈
                stack.pop()

            else:
                #如果得到了一个左括号，就入栈
                stack.append(s[i]) 
            i+=1
                


        #栈中无数即可
        return len(stack)==0
