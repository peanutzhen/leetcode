#include <stdio.h>

int
main()
{
    int n;
    int max_len = 1, len = 1;
    int prev, cur;

    scanf("%d", &n);
    scanf("%d", &prev);
	for(int i = 1; i < n; i++)
    {
        scanf("%d", &cur);
		if(cur >= prev)
        {
			len++;
			max_len = (len > max_len) ? len : max_len;
		} else {
            len = 1;
		}
		prev = cur;
	}

    printf("%d\n", max_len);
    return 0;
}