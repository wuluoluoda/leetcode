class Solution(object):
    def twoSum(self, nums, target):
        m = {}
        for index, val in nums :
            if  m[target-val] :
                return {m[target-val], index}
            else :
                m[val] = index
            
        
        return {}