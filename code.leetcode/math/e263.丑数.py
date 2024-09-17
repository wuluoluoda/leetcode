class Solution(object):
    def isUgly(self, n):
        if n == 1 :
            return True
	

        if n <= 0 :
            return False
        

        while n != 1 :
            bit = 0
            s2 = n % 2
            if s2 == 0 :
                n = n / 2
                bit+=1
            
            s3 = n % 3
            if s3 == 0 :
                n = n / 3
                bit+=1
            
            s5 = n % 5
            if s5 == 0 :
                n = n / 5
                bit+=1
            
            if bit == 0 :
                return False
            
        
        return True