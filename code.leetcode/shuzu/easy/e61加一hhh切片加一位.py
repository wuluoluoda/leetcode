class Solution(object):
    def plusOne(self, digits):
        n = len(digits)
        digits[n-1]+=1
        #对数组处理，进位处理
        #如何进位，首先从后往前遍历，直到某位不为9
        right = n - 1
        while  right >= 0 :
            if digits[right] != 10 :
                break
            if right-1 >= 0 :
                digits[right] = 0
                digits[right-1]+=1
            else :
                digits[right] = 0
                #首位加一位
                digits.insert(0,1)            
            right-=1
            
        
        return digits