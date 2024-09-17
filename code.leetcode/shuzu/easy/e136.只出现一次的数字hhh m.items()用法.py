class Solution(object):
    def singleNumber(self, nums):
        m = {}
        #if 存在，反转bool
        #不存在，存入，下一个
        for  v in nums :
            if v in m  :
                m[v] = not m[v]
            else :
                m[v] = True
            
        
        #遍历map，返回true的值
        for k, v in m.items():
            if v :
                return k
            
        
        return 0
