#include <stdio.h>
#include <stdlib.h>

double func(int k,int *nums1, int nums1Size, int *nums2, int nums2Size){
    if(nums1Size == 0)
        return nums2[k];
    if(nums2Size == 0)
        return nums1[k];
    int low[2]={0,0};
    int mid[2]={(nums1Size-1)/2,(nums2Size-1)/2};
    int high[2]={nums1Size-1,nums2Size-1};
    while(1)
    {
        if(low[0]>high[0]){
            return (double)nums2[k-low[0]];
        }
        else{
            // 判断插入点
            int index = mid[0];
            if(high[1]<low[1]) // 说明已经找到插入点，该点前有low[1]个值
                index += low[1];
            else{
                if(nums1[mid[0]] <= nums2[mid[1]]){
                    high[1] = mid[1]-1;
                    mid[1] = (low[1]+high[1]) / 2;
                    continue;
                }
                else{
                    low[1] = mid[1]+1;
                    mid[1] = (low[1]+high[1]) / 2;
                    continue;
                }
            }

            // 比较是否为第k+1小的值
            if(index < k){ //小了
                low[0]=mid[0]+1;
                mid[0]=(low[0]+high[0])/2;
                //clear nums2 counter.
                low[1]=0;
                high[1]=nums2Size-1;
                mid[1]=(low[1]+high[1]) / 2;
                continue;
            }
            else if(index > k){ //大了
                high[0]=mid[0]-1;
                mid[0]=(low[0]+high[0])/2;
                //clear nums2 counter.
                low[1]=0;
                high[1]=nums2Size-1;
                mid[1]=(low[1]+high[1]) / 2;
                continue;
            }
            else{ //刚好fit
                return (double)nums1[mid[0]];
            }
        }
    }
}

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int total_length = nums1Size + nums2Size;
    if(total_length % 2 == 0)
        return 0.5 * (func(total_length/2,nums1,nums1Size,nums2,nums2Size) + func((total_length-1)/2,nums1,nums1Size,nums2,nums2Size));
    else
        return func(total_length/2,nums1,nums1Size,nums2,nums2Size);
}

int main(){
    FILE *input_file = fopen("test/001.txt","r");
    int *ar1, ar1Size;
    int *ar2, ar2Size;
    double ans;

    while(1){
        if(feof(input_file))
            break;

        fscanf(input_file,"%d",&ar1Size);
        ar1 = (int *)malloc(ar1Size*sizeof(int));
        for(int i=0;i<ar1Size;i++)
            fscanf(input_file,"%d",ar1+i);

        fscanf(input_file,"%d",&ar2Size);
        ar2 = (int *)malloc(ar2Size*sizeof(int));
        for(int i=0;i<ar2Size;i++)
            fscanf(input_file,"%d",ar2+i);

        fscanf(input_file,"%lf",&ans);

        if(ans - findMedianSortedArrays(ar1,ar1Size,ar2,ar2Size) < 0.01)
            putchar('.');
        else
            putchar('x');

        free(ar1);
        free(ar2);
    }

    putchar('\n');
    fclose(input_file);
    return 0;
}