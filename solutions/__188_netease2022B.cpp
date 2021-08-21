/*
 *      Author: peanutzhen
 *      Created time: 2021/8/21 15:02:35
 */

#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    /**
     * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
     *
     * 返回Sn的第k位字符
     * @param n int整型 Sn的n
     * @param k int整型 需要返回的字符下标位
     * @return char字符型
     */
    char findKthBit(int n, int k) {
        // write code here
        int len = (int)pow(2,n)-1;
        char s[len];
        char tmp[len];
        memset(s, '\0', sizeof(s));
        memset(tmp, '\0', sizeof(s));
        s[0] = 'a';
        int ptr = 1; // 指向si+1将要写入的位置
        for (int i = 2; i <= n; ++i) {
            s[ptr] = (char)('a' + i - 1);
            // reverse
            reverse_copy(s, s+ptr, tmp);
            // invert
            for (int j = 0; j < ptr; ++j) {
                tmp[j] = (char)('a' + 25 - (tmp[j] - 'a'));
            }
            int p = ptr;
            for (int j = 0; j < ptr; ++j) {
                s[++p] = tmp[j];
            }
            ptr = p + 1;
        }
        return s[k-1];
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    auto* solution = new Solution();
    cout << solution->findKthBit(4,11);
    return 0;
}

