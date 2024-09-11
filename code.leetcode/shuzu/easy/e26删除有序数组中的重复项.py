class Solution(object):
    def removeDuplicates(self, nums):
        #设map
        m = {}
        #遍历数组
        left = 0
        for  v in nums:  # v 即 nums[right]
            if v not in m:
                nums[left] = v
                left+=1
                m[v]=True
            
        
        
        return left
        
	