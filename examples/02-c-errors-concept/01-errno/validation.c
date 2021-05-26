#include <stdlib.h>
#include <stdio.h>
#include <string.h>

typedef enum {
    VALID_ERR_OK = 0,
    VALID_ERR_INVALID_USERNAME,
    VALID_ERR_INVALID_EMAIL,
    VALID_ERR_WEAK_PASSWORD,
    VALID_ERR_COUNT, // Служебное поле для определения размера enum.
} validation_error_t;

const char* const VALIDATION_ERROR_STRS[] = {
    "All is OK",
    "Invalid username",
    "Invalid email",
    "Too weak password",
};

const char* validation_error_str(validation_error_t err)
{
    if (VALID_ERR_OK <= err && err < VALID_ERR_COUNT) {
         return VALIDATION_ERROR_STRS[err];
    }
    return "Unknown";
}

const int PASS_MIN_LEN = 10;

typedef struct {
    char *username;
    char *email;
    char *password;
} user_t;

validation_error_t validate(user_t u)
{
    if (strlen(u.password) < PASS_MIN_LEN) {
        return VALID_ERR_WEAK_PASSWORD;
    }
    // ...
    return 0;
}

int main()
{
    validation_error_t err = validate((user_t){"bob", "bob@gmail.com", "bob123"});
    if (err != 0) {
         printf("user validation err: %s", validation_error_str(err));
    }

    return 0;
}
