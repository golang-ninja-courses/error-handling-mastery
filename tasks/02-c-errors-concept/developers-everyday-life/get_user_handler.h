#ifndef GET_USER_HANDLER_H
#define GET_USER_HANDLER_H

#include <stdlib.h>
#include "db.h"
#include "marshalers.h"

typedef enum {
    HTTP_ERR_OK = 0, // 200
    // Реализуй нас.
    HTTP_ERR_BAD_REQUEST = 400,
    HTTP_ERR_UNPROCESSABLE_ENTITY = 422,
    HTTP_ERR_NOT_FOUND = 404,
    HTTP_ERR_INTERNAL_SERVER_ERROR = 500,
} http_error_t;

const char* const HTTP_ERR_STRS[] = {
    "200 OK",
    // Реализуй нас.
    // ...
    "400 Bad Request",
    "422 Unprocessable Entity",
    "404 Not Found",
    "500 Internal Server Error",
};

const char *http_error_str(http_error_t err)
{
    return HTTP_ERR_STRS[err];
}

http_error_t get_user_handler(char *request_data, char **response_data)
{
    http_error_t err = HTTP_ERR_OK;
    // Реализуй меня.
    
    int target_user_id = 0;
    request_t *req = NULL;
    target_user_id = unmarshal_request(request_data, &req);
    if (target_user_id == -1) {
        err = HTTP_ERR_BAD_REQUEST;
        //free(req->user_id);
        free(req);
        return err;
    };
    if (target_user_id <= 0) {
        err = HTTP_ERR_UNPROCESSABLE_ENTITY;
        //free(req->user_id);
        free(req);
        return err;
    }

    user_t *target_user_struct = NULL;
    db_error_t db_err = 0;
    db_err = db_get_user_by_id(target_user_id, &target_user_struct);
    if (db_err == DB_ERR_NOT_FOUND) {
        err = HTTP_ERR_NOT_FOUND;
        free(target_user_struct);
        return err;
    }
    if (db_err == DB_ERR_INTERNAL) {
        err = HTTP_ERR_INTERNAL_SERVER_ERROR;
        free(target_user_struct);
        return err;
    }

    int resp_int = 0;
    resp_int = marshal_response((response_t){target_user_struct}, response_data);
    if (resp_int == -1 ) {
        err = HTTP_ERR_INTERNAL_SERVER_ERROR;
        //free(resp);
        return err;
    }

    return err;
}

#endif
