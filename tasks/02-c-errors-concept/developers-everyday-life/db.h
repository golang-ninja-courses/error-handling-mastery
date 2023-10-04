#ifndef DB_H
#define DB_H

#include <stdlib.h>
#include <string.h>

#define KNOWN_USER 4224
#define INTERNAL_ERR_THRESHOLD 10000

typedef struct {
    int  id;
    char *email;
} user_t;

typedef enum {
    DB_ERR_OK        = 0,
    DB_ERR_INTERNAL  = 1,
    DB_ERR_NOT_FOUND = 2,
} db_error_t;

db_error_t db_get_user_by_id(int uid, user_t **user)
{
    *user = NULL;

    if (uid >= INTERNAL_ERR_THRESHOLD) {
        return DB_ERR_INTERNAL;
    }

    if (uid != KNOWN_USER) {
        return DB_ERR_NOT_FOUND;
    }

    user_t *u = (user_t *) malloc(sizeof(user_t));
    if (u == NULL) {
        return DB_ERR_INTERNAL;
    }

    u->id = uid;

    u->email = (char *) calloc(14, sizeof(char));
    strncpy(u->email, "bob@gmail.com\0", 14);
    if (u->email == NULL) {
        free(u);
        return DB_ERR_INTERNAL;
    }

    *user = u;
    return DB_ERR_OK;
}

#endif
