class Solution(object):
    def lengthOfLastWord(self, s):
        i = len(s) - 1
        while i >= 0:
            if s[i] != ' ' :
                j = i
                while j >= 0:
                    if s[j] == ' '  :
                        return i - j
                    if j == 0 :
                        return i + 1
                    j -= 1
            i-=1		
			
		
		

        return 0