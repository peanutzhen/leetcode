#include<stdio.h>
#include<string.h>
#include<sys/time.h>

#define MAX_COLOR 100
int memo[MAX_COLOR][MAX_COLOR][MAX_COLOR]={{{0}}};

int removeBoxes(int* boxes, int boxesSize){
    // 初始化基准情形
    for(int i=0;i<boxesSize;i++){
        for(int perfix=0;perfix<=i;perfix++){
            // 从i开始至i结束，且boxes[i]在i之前已经出现过perfix次。
            // 显然，perfix不可能出现i+1次以上
            memo[i][i][perfix] = (1+perfix) * (1+perfix);
        }
    }
    // 自底向上（求解整个解空间）
    for(int len=2;len<=boxesSize;len++){
        for(int start=0;start+len-1<boxesSize;start++){
            for(int perfix=0;perfix<=start;perfix++){
                // 简单地移除首项
                int max_value= (1+perfix)*(1+perfix) + memo[start+1][start+len-1][0];
                // 考虑合并
                for(int merge=start+1;merge<=start+len-1;merge++){
                    if(boxes[merge] == boxes[start]){
                        if(max_value < memo[start+1][merge-1][0]+memo[merge][start+len-1][1+perfix])
                            max_value = memo[start+1][merge-1][0]+memo[merge][start+len-1][1+perfix];
                    }
                }
                memo[start][start+len-1][perfix]=max_value;
            }
        }
    }
    return memo[0][boxesSize-1][0];
}
int main(){
    int ar1[9] = {1, 3, 2, 2, 2, 3, 4, 3, 1};
    int ar2[21] = {4, 3, 2, 1, 2, 3, 2, 1, 2, 3, 4, 3, 2, 1, 2, 3, 2, 1, 2, 3, 4};
    int ar3[4]={1,2,3,4};
    int ar4[]={20,12,2,11,6,18,4,6,8,12,16,18,15,6,10,8,20,8,15,16};
    struct timeval start,end;
    gettimeofday(&start,NULL);
    printf("%d\n", removeBoxes(ar1,sizeof(ar1)/sizeof(int)));
    //printf("%d\n", removeBoxes(ar3,sizeof(ar3)/sizeof(int)));
    //printf("Ans:%d\n", removeBoxes(ar2,sizeof(ar2)/sizeof(int)));
    //printf("Ans:%d\n", removeBoxes(ar4,sizeof(ar4)/sizeof(int)));
    gettimeofday(&end,NULL);
    long diff = 1000000 * (end.tv_sec-start.tv_sec)+ end.tv_usec-start.tv_usec;
    printf("自顶向下：%ld us\n",diff);
    return 0;
}