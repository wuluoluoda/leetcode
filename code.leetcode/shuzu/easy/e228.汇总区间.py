class Solution(object):
    def summaryRanges(self, nums):
        #go代码分组循环法示例
        """func summaryRanges(nums []int) (ans []string) :
    for i, n := 0, len(nums); i < n; :
        left := i
        for i++; i < n && nums[i-1]+1 == nums[i]; i++ :
        
        s := strconv.Itoa(nums[left])
        if left < i-1 :
            s += "->" + strconv.Itoa(nums[i-1])
        
        ans = append(ans, s)
    
    return



        """
        ans = []
        i, n = 0,len(nums)
        while i < n :
            left = i
            i+=1
            while i < n and nums[i-1]+1 == nums[i]:
                i+=1
                
            s = str(nums[left])
            if left < i-1 :
                s += "->" + str(nums[i-1])
            
            ans.append(s)

        
        return ans