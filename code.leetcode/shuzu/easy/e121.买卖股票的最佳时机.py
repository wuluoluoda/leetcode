class Solution(object):
    # max 返回两个整数中较大的一个
    def max(self,a, b): 
        if a > b :
            return a
        
        return b


    # min 返回两个整数中较小的一个
    def min(self,a, b):
        if a < b :
            return a
    
        return b
    def maxProfit(self, prices):
       
        inf = 10000
        minPrice = inf
        maxProfit = 0

        for  price in  prices :
            maxProfit = max(maxProfit, price-minPrice)
            minPrice = min(minPrice, price)
        
        return maxProfit


