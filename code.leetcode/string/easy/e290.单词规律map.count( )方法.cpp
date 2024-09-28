//byte不是c++基本类型,char才是***************
#include <vector>
class Solution {
public:
    bool wordPattern(string pattern, string s) {
 //用空格分裂字符串s，存入words
 std::vector<std::string> words;
std::istringstream iss(s);
std::string word;
while (iss >> word) {
    words.push_back(word);
}
	//长度不等直接返回false
	if (pattern.size() != words.size()) {
		return false;
	}
	unordered_map<char,string>patternMap = {};
	unordered_map<string,char>wordsMap = {};
	
	for (int i =0;i<pattern.size();i++) {
		//存在值但不对应就返回false
		if (patternMap.count(pattern[i])  && patternMap[pattern[i]] != words[i]) {
			return false;

		}
		if (wordsMap.count(words[i]) && wordsMap[words[i]] != pattern[i]) {
			return false;
		}
		patternMap[pattern[i]] = words[i];
		wordsMap[words[i]] = pattern[i];
	}
	return true;
    }
};