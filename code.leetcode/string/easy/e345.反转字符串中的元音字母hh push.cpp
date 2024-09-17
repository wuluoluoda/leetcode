#include <string>
class Solution {
public:
    string reverseVowels(string s) {
/*ss = []rune(s)*/
vector<char> ss(s.begin(), s.end()); // 将字符串转换为字符向量
    /*std::vector<char>ss;
    for(auto c : s) ss.push_back(c);*/
	/*unordered_map<char,bool>vowels = {
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	};*/
	unordered_map<char, bool> vowels = {
            {'a', true},
            {'e', true},
            {'i', true},
            {'o', true},
            {'u', true},
            {'A', true},
            {'E', true},
            {'I', true},
            {'O', true},
            {'U', true},
        };
	//如果存在元音字母，则记录其位置
	vector<int>pos;
	
	/*for (auto i,auto c : s) {
		if (vowels.count(c)) {
			//pos.push.back(i)=pos = append(pos, i)***********************************************
			pos.push_back(i);
		}
			
		}*/
	for (size_t i = 0; i < s.size(); ++i) {
            if (vowels.count(s[i])) {
                pos.push_back(i);
            }
        }
	
	//对记录的位置进行反转
	//i < pos.size()-1-i在c++中是错误的，应该是i < pos.size()/2
	/*for (int i = 0; i < pos.size()-1-i; ++i) {
		swap(ss[pos[i]], ss[pos[pos.size()-1-i]]);
	}*/
	// 对记录的位置进行反转
       for (size_t i = 0; i < pos.size()/2; ++i) {
            swap(ss[pos[i]], ss[pos[pos.size()-1-i]]);
        }
	/*return string(ss)*/
	return string(ss.begin(), ss.end()); // 将字符向量转换回字符串
    }
};