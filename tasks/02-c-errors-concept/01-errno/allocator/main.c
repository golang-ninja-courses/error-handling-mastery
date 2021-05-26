// gcc main.c -o main
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "allocator.h"

int main()
{
	size_t size = 0;
	int uid = 0;

	if (scanf("%d %zd", &uid, &size) == 0) {
		perror("scanf failed");
		exit(1);
	}

	void *p = allocate(uid, size);
	if (p == NULL) {
		printf("allocation error: %s", strerror(errno));
		exit(1);
	}

	printf("allocation was successful for %zu bytes", size);
	free(p);

    return 0;
}
