#include <stdio.h>
#include <errno.h>

extern int errno;

int main()
{
    FILE *fp;
    fp = fopen("unexistent_file.txt", "r");

    printf("value of errno: %d\n", errno); // value of errno: 2
    return 0;
}
