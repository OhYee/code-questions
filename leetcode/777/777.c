#include <stdbool.h>

bool canTransform(char *start, char *end) {
    char *a = start;
    char *b = end;
    while (*a != '\0') {
        // printf("%c %d\n", *a, *a);
        if (*a == 'R' && *b == 'X') {
            char *temp = a;
            while (*temp != '\0') {
                ++temp;
                if (*temp == 'L')
                    return false;
                if (*temp == 'X') {
                    *temp = 'R';
                    break;
                }
            }
            if (*temp == '\0')
                return false;
            *a = 'X';
        } else if (*a == 'X' && *b == 'L') {
            char *temp = a;
            while (*temp != '\0') {
                ++temp;
                if (*temp == 'R')
                    return false;
                if (*temp == 'L') {
                    *temp = 'X';
                    break;
                }
            }
            if (*temp == '\0')
                return false;
            *a = 'L';
        } else if (*a != *b)
            return false;
        ++a;
        ++b;
    }
    return true;
}