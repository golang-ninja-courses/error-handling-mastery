#include <stdio.h>
#include <errno.h>

extern int errno;

int main()
{
    errno = 0;
    perror("msg for zero errno\t");

    errno = 4242;
    perror("msg for custom errno\t");

    return 0;
}
