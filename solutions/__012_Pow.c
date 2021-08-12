#include <stdio.h>
#include <math.h>

double myPow(double x, int n){
    // 减轻时间
    if(x == 1)
        return 1;

    if(n == 0)
        return 1;
    if(n < 0)
    {
        if(n == -2147483648)    // 避免溢出
            return myPow(1/x,2147483647) * 1/x;
        return myPow(1/x,-n);
    }
    else
    {
        if(n % 2 == 1)
        {
            return myPow(x,n-1) * x;
        }
        else
        {
            return myPow(x,n/2) * myPow(x,n/2); // 二分，减少运算次数
        }
    }
}

int
main()
{
    double x = -1;
    int n = -31;
    printf("pow: %f\n",pow(x,n));
    printf("myPow: %f\n",myPow(x,n));
    printf("%d",-31 % 2);
    return 0;
}