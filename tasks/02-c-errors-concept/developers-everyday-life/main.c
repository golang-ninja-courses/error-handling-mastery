#include <stdio.h>
#include <stdlib.h>

#include "get_user_handler.h"

int main()
{
    const int n = 100;
    char request_data[n];

    if (fgets(request_data, n, stdin) == NULL) {
        return EXIT_FAILURE;
    }

    char *resp = NULL;
    http_error_t err = get_user_handler(request_data, &resp);

    puts(http_error_str(err));

    if (err) {
        if (resp != NULL) {
            puts("Response is not NULL after http error!");
            return EXIT_FAILURE;
        }

    } else {
        puts(resp);
        free(resp);
    }

    return EXIT_SUCCESS;
}
