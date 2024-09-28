class Solution(object):
    def containsDuplicate(self, nums):
        #建map
        m={}
        #遍历
        for v in nums:
            #存在，返回

            #不存在，添加
            if v in m:
                return True
            
            m[v]=True
        	
        return False
            #存在，返回

            #不存在，添加