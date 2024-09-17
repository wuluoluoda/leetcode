class Solution(object):
    def missingNumber(self, nums):
        has = {}
        for  v in nums :
            has[v] = True
        i = 0
        while True:
            if i not in has :
                return i
            i+=1
        