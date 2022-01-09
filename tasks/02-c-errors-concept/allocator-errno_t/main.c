#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "allocator.h"

int main()
{
    size_t size = 0;
    int uid = 0;

    errno = 0;
    if (scanf("%d %zu", &uid, &size) != 2) {
        perror("scanf failed");
        exit(1);
    }

    void *p = NULL;
    errno_t err = allocate(uid, size, &p);
    if (err != 0) {
        printf("allocation error: %s\n", strerror(err));
        exit(0); // Считаем валидной ситуацией.
    }
    if (p == NULL) {
        printf("memory pointer is NULL after allocation");
        exit(1);
    }

    printf("allocation was successful for %zu bytes\n", size);
    free(p);

    return 0;
}
