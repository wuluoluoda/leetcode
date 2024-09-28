class Solution(object)
    def merge(self, nums1, m, nums2, n)
        #遍历nums2放入nums1的n部分
        i=0
        while i<n:
            nums1[m+i]=nums2[i]
            i+=1  
        
        #排序
        nums1.sort()
