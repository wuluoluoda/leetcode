class Solution(object):
    def addDigits(self, num):
        #遍历
        
        while num >= 10 :
            sum = 0
            while num > 0 :
                sum += num % 10
                num /= 10
            
            num = sum
        
        return num