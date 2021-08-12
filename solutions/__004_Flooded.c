#include <stdio.h>
#include <stdlib.h>
typedef struct info{
    int row;
    int col;
    int *elevations;
    int rainTotal;
    struct info *next;
}*Info;
typedef struct queue{
    Info front;
    Info rear;
    int size;
}*Queue;
void addRegion(Queue q, int r,int c){
    if(q->size==0)
    {
        q->front=(Info)malloc(sizeof(struct info));
        q->rear=q->front;
        q->rear->next=NULL;
        q->rear->row=r;
        q->rear->col=c;
        q->rear->elevations=(int *)malloc(r*c*sizeof(int));
        for(int i=0;i<r*c;i++)
            scanf("%d",q->rear->elevations+i);
        scanf("%d",&q->rear->rainTotal);
    }
    else{
        q->rear->next=(Info)malloc(sizeof(struct info));
        q->rear=q->rear->next;
        q->rear->next=NULL;
        q->rear->row=r;
        q->rear->col=c;
        q->rear->elevations=(int *)malloc(r*c*sizeof(int));
        for(int i=0;i<r*c;i++)
            scanf("%d",q->rear->elevations+i);
        scanf("%d",&q->rear->rainTotal);
    }
    q->size++;
}
void insertion_sort(int *arr, int len){
    int i,j,temp;
    for (i=1;i<len;i++){
        temp = arr[i];
        for (j=i;j>0 && arr[j-1]>temp;j--)
            arr[j] = arr[j-1];
        arr[j] = temp;
    }
}
int main() {
    Queue q=(Queue)malloc(sizeof(struct queue));
    q->front=NULL;
    q->rear=NULL;
    q->size=0;

    int rows,cols;
    scanf("%d%d",&rows,&cols);

    if(rows!=0&&cols!=0)
        addRegion(q,rows,cols);

    Info tmp=q->front;
    for(int region=0;region<q->size; region++){
        insertion_sort(tmp->elevations,(tmp->row)*(tmp->col));
        int sum=0;
        int i;
        for(i=0;i<(tmp->row)*(tmp->col);i++){
            int j=i-1;
            if(j!=-1)
                sum+=tmp->elevations[j];
            if(100*i*(tmp->elevations[i])-100*sum<tmp->rainTotal)
                continue;
            else
                break;
        }
        if(tmp->rainTotal<=0)
        {
            printf("Region %d\n",region+1);
            if(tmp->rainTotal==0)
                printf("Water level is %.2f meters.\n",(float)tmp->elevations[0]);
            else
                printf("Water level is -inf meters.\n");
            printf("0.00 percent of the region is under water.\n\n");
            scanf("%d%d",&rows,&cols);
            if(rows!=0&&cols!=0)
                addRegion(q,rows,cols);
            tmp=tmp->next;
            continue;
        }
        float height;
        float percent;
        if(i==(tmp->row)*(tmp->col)){
            sum+=tmp->elevations[i-1];
            height=(float)(tmp->rainTotal+100*(sum))/(float)(100*i);
            percent=100.00f;
        }
        else{
            height=(float)(tmp->rainTotal+100*(sum))/(float)(100*i);
            percent=(float)(i*100)/(float)((tmp->row)*(tmp->col));
        }

        printf("Region %d\n",region+1);
        printf("Water level is %.2f meters.\n",height);
        printf("%.2f percent of the region is under water.\n\n",percent);
        scanf("%d%d",&rows,&cols);
        if(rows!=0&&cols!=0)
            addRegion(q,rows,cols);
        tmp=tmp->next;
    }
    return 0;
}
