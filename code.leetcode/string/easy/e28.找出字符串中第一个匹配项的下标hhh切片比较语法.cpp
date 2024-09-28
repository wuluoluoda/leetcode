class Solution {
public:
    int strStr(string haystack, string needle) {
//needle更长，直接返回
if (needle.size() > haystack.size()){
    return -1;
}
	//ha

	//计算needle的长度1234  345
//把needle反向放入栈中
//遍历haystack，如果haystack[i] == needle[0]，则从i开始，比较haystack[i:i+len(needle)]和needle是否相等
	
	for (int i=0;i<haystack.size()-needle.size()+1;++i){
    if (haystack[i]== needle[0]) {
       // 使用 std::equal 来比较子字符串
                if (std::equal(needle.begin(), needle.end(), haystack.begin() + i)){
           return i;
       }
    }
}
return -1;
    }
};