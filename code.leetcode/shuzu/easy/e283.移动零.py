class Solution(object):
    def moveZeroes(self, nums):
                #先删掉再补全
            #双指针
        left=0
        i=0
        while i<len(nums):
            if nums[i]!=0:
                nums[left]=nums[i]
                left+=1
            i+=1
            

        #补全
        j=left
        while j<len(nums):
            nums[j]=0
            j+=1
