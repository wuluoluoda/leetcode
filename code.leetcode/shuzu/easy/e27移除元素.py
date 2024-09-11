class Solution(object):
    def removeElement(self, nums, val):
        left = 0
        for  v in nums :
            if v != val :
                nums[left] = v
                left+=1
            
        
        return left