#ifndef GET_USER_HANDLER_H
#define GET_USER_HANDLER_H

#include <stdlib.h>
#include "db.h"
#include "marshalers.h"

typedef enum {
    HTTP_ERR_OK = 0, // 200
    // Реализуй нас.
    // ...
} http_error_t;

const char* const HTTP_ERR_STRS[] = {
    "200 OK",
    // Реализуй нас.
    // ...
};

const char *http_error_str(http_error_t err)
{
    return HTTP_ERR_STRS[err];
}

http_error_t get_user_handler(char *request_data, char **response_data)
{
    http_error_t err = HTTP_ERR_OK;
    // Реализуй меня.
    // ...
    return err;
}

#endif
