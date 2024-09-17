class Solution(object):
    def isPowerOfTwo(self, n):
        big = 1 << 30
        return n > 0 and big%n == 0