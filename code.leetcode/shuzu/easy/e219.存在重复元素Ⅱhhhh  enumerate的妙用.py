class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        #建map
        m={}
        #遍历
        #enumerate的用处
        for i, v in enumerate(nums):
            #存在，返回

            #不存在，添加
            if v in m and i-m[v]<=k:
                return True
            
            m[v]=i
        	
        return False
            #存在，返回

            #不存在，添加