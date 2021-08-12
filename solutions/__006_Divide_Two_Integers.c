#include <stdio.h>
/*
 * 非常暴力的做法，迭代，就硬迭代
 * 并作结果溢出检测
 * 简单情形直接ifelse解决，避免Time exceed
 */

int
divide(int dividend, int divisor)
{
    // 简单情形处理：避免时间溢出
    if(divisor == -1)
        if(dividend == -2147483648)
            return 2147483647;
        else
            return -dividend;

    if(divisor == 1)
        return dividend;

    // 我知道下面的代码很糟糕，但过了就行！！
    int rtv = 0;
    if(dividend < 0 && divisor < 0)
    {
        while(dividend <= divisor)
        {
            dividend -= divisor;
            if(rtv == 2147483647)
                return 2147483647;
            rtv += 1;
        }
    }
    else if(dividend < 0 && divisor > 0)
    {
        while(dividend + divisor <= 0)
        {
            dividend += divisor;
            if(rtv == -2147483648)
                return 2147483647;
            rtv -= 1;
        }
    }
    else if((dividend > 0 && divisor < 0))
    {
        while(dividend + divisor >= 0)
        {
            dividend += divisor;
            if(rtv == -2147483648)
                return 2147483647;
            rtv -= 1;
        }
    }
    else
    {
        while(dividend >= divisor)
        {
            dividend -= divisor;
            if(rtv == 2147483647)
                return 2147483647;
            rtv += 1;
        }
    }
    return rtv;
}

int
main()
{
    printf("%d\n", divide(-2147483648,-1));
    return 0;
}