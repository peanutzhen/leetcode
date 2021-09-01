/*
 *      Author: peanutzhen
 *      Created time: 2021/9/1 19:04:48
 */

#include <bits/stdc++.h>

using namespace std;

string removeKdigits(string num, int k) {
    // 单调栈 元素严格递增
    vector<char> mono_stack;
    for (auto digit: num) {
        while (!mono_stack.empty() && mono_stack.back() > digit && k) {
            mono_stack.pop_back();
            --k; // 移除一次
        }
        mono_stack.push_back(digit);
    }
    // 注意，移除次数不到k次(那么，栈的单调性仍能满足)
    while (k--)
        mono_stack.pop_back();
    string ans;
    for (auto digit: mono_stack) {
        // 注意，栈中可能以0为前导
        if (digit != '0' || !ans.empty())
            ans += digit;
    }
    return ans.empty() ? "0" : ans;
}

int main() {

    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    cout << removeKdigits("9", 1) << "\n";
    return 0;
}

