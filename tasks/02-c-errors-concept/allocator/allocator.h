#include <errno.h>

extern int errno;

#define ADMIN 777
#define MIN_MEMORY_BLOCK 1024

void *allocate(int user_id, size_t size)
{
    if (user_id != ADMIN) {
        errno = 1;
        return NULL;
    }
    if (MIN_MEMORY_BLOCK > size ) {
        errno = 33;
        return NULL;
    }
    void *m = malloc(size);
    if (m == NULL) {
        errno = 12;
    }
    return m;
}
