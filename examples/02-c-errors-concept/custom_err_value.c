#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>

typedef int errno_t;

const errno_t ESOMETHINGREALLYBAD = 4242;

errno_t g()
{
    // ...
    int something_really_bad_happens = 1;

    // ...
    if (something_really_bad_happens == 1) {
        return ESOMETHINGREALLYBAD;
    }

    // ...
    return 0;
}

int main()
{
    errno_t err = g();
    if (err != 0) {
        puts(strerror(err)); // Unknown error: 4242
        return EXIT_FAILURE;
    }

    return EXIT_SUCCESS;
}
