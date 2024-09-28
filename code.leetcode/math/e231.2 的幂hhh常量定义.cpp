class Solution
{
public:
    bool isPowerOfTwo(int n)
    {
        static constexpr int big = 1 << 30;
        return n > 0 && big % n == 0
    }
};