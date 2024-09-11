class Solution
{
public:
    int lengthOfLastWord(string s)
    {
        for (int i = s.size() - 1; i >= 0; --i)
        {
            if (s[i] != ' ')
            {
                for (int j = i; j >= 0; --j)
                {
                    if (s[j] == ' ')
                    {
                        return i - j;
                    }
                    else if (j == 0)
                    {
                        return i + 1;
                    }
                }
            }
        }
        return 0;
    }
};