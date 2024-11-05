#include "exercise.h"
#include <stdio.h>

int main(void) {
        // snek_int_t integer;
        // integer.value = 42;
        // snek_zero_out(&integer, INTEGER);
        int a = 1;
        int b = 2;
        printf("a: %d\n", a);
        printf("b: %d\n", b);
        int tmp = a;
        a = b;
        b = tmp;
        printf("a: %d\n", a);
        printf("b: %d\n", b);
}
