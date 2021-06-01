#include <errno.h>

#define __STDC_WANT_LIB_EXT1__ 1
// Если define выше не работает для нашего компилятора, то определяем тип руками:
// typedef int errno_t;

extern int errno;

#define ADMIN 777
#define MIN_MEMORY_BLOCK 1024

errno_t allocate(int user_id, size_t size, void **mem)
{
    // Реализуй меня.
}
