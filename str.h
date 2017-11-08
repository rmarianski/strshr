#ifndef STR_H
#define STR_H

typedef struct {
    unsigned int len;
    char *str;
} string_result;

string_result make_string();

#endif
