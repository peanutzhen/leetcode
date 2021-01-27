#include <stdio.h>

// 一开始时间溢出，加个记录表就可以避免
// 因为子问题有重复的部分
int table[101][101] = {{0}};

int uniquePaths(int m, int n){
    if(m == 1 || n == 1)
        return 1;
    if(table[m][n] == 0)
        table[m][n] = uniquePaths(m-1,n) + uniquePaths(m,n-1);
    return table[m][n];
}

int
main()
{
    printf("%d\n", uniquePaths(23,12));
    return 0;
}