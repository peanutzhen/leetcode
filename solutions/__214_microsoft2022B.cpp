/*
 *      Author: peanutzhen
 *      Created time: 2021/9/21 15:26:37
 */

#include <bits/stdc++.h>

using namespace std;

// 题目描述：
// 给定用字符串描述的整数S
// 你可以至多1次在任意位上将数字改变成(0~9)
// 求这样操作得到能被3整除的数有多少？
// S in [1, 100000]，且S可以有前导0，即"0023"这样
// IO example:
// S = "23" -> 7
// reason: "03", "33", "63", "93", "21", "24", "27"
int solution(string &S) {
    int ans = (stoi(S) % 3 == 0) ? 1 : 0;
    int length = (int) S.size();
    for (int i = 0; i < length; ++i) {
        char ori = S[i];
        for (int j = 0; j < 10; ++j) {
            S[i] = char('0' + j);
            if (ori == S[i])
                continue;
            int num = stoi(S);
            if (num % 3 == 0) {
//                cout << num << "\n";
                ++ans;
            }
        }
        S[i] = ori;
    }
    return ans;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    string S = "23";
    int ans = solution(S);
    cout << "ans: " << ans << endl;
    return 0;
}

