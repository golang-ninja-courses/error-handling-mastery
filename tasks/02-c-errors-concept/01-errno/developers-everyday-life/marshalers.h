#ifndef MARSHALERS_H
#define MARSHALERS_H

#include <stdio.h>
#include <stdlib.h>
#include "db.h"

typedef struct {
    int user_id;
} request_t;

int unmarshal_request(char *request_data, request_t **request)
{
    *request = NULL;

    request_t *r = (request_t *) malloc(sizeof(request_t));
    if (r == NULL) {
        return -1;
    }

    if (sscanf(request_data, "{\"user_id\": %d}", &(r->user_id)) == 0) {
        free(r);
        return -1;
    }

    *request = r;
    return 0;
}

typedef struct {
    user_t *user;
} response_t;

int marshal_response(response_t response, char **response_data)
{
    char *buf = (char *) calloc(100, sizeof(char));

    if (sprintf(buf, "{\"user\": {\"id\": \"%d\", \"email\": \"%s\"}", response.user->id, response.user->email) == 0) {
        free(buf);
        *response_data = NULL;
        return -1;
    }

    *response_data = buf;
    return 0;
}

#endif
