#include <stdio.h>
#include <errno.h>

int main()
{
    FILE *fp;
    fp = fopen("unexistent_file.txt", "r");

    perror("cannot open file"); // cannot open file: No such file or directory
    return 0;
}
