class Solution(object):
	def isPalindrome(self, x) :
		num=0 
		biw = x
		#读出正负数

		if x >= 0 : 
			#123%10=3 123/10=12 12%10=2 12/10=1 1%10=1 1/10=0
			while biw > 0 :

				num = num*10 + biw%10
				biw = biw / 10
		
	
		if num == x :
			return True
	
		return False
