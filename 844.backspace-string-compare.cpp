#include <string>
#include <iostream>
using namespace std;
/*
 * @lc app=leetcode id=844 lang=cpp
 *
 * [844] Backspace String Compare
 */

// @lc code=start
class Solution {
public:
    bool backspaceCompare(string s, string t) {
        int i = 0;
        int k = 0;
        for (int j=0; j<s.size(); j++) {
            if (s[j] == '#') {
                if (i>0) {
                    i--;
                }
            } else {
                s[i] = s[j];
                i++;
            }
        }

        for (int j=0; j<t.size(); j++) {
            if (t[j] == '#'){
                if (k>0) {
                    k--;
                }
            }
            if (t[j] != '#') {
                t[k] = t[j];
                k++;
            }
        }

        if (i == k) {
            for (int j=0; j<i; j++){
                if (s[j] != t[j]){
                    return false;
                }
            }
            return true;
        }
        return false;
    }
};
// @lc code=end
