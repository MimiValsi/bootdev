#include "exercise.h"
<<<<<<< HEAD

int main(void) {}
=======
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void swap(int *a, int *b);
void swap_s(char **a, char **b);
void swap_g(void *vp1, void *vp2, size_t size);

int main(void) {}

stack_t *stack_new(size_t capacity) {
        stack_t *stack = malloc(capacity);
        if (!stack) {
                return NULL;
        }
        stack->count = 0;
        stack->capacity = capacity;
        stack->data = malloc(capacity * sizeof(void *));
}

void swap_g(void *vp1, void *vp2, size_t size) {
        void *tmp = malloc(size * sizeof(void *));
        if (!tmp) {
                return;
        }

        memcpy(tmp, vp2, size);
        memcpy(vp2, vp1, size);
        memcpy(vp1, tmp, size);
}

void swap_s(char **a, char **b) {
        char *tmp = *a;
        *a = *b;
        *b = tmp;
}

void swap(int *a, int *b) {
        int tmp = *a;
        *a = *b;
        b = &tmp;
}
>>>>>>> 84524d623211a705e578a3e72876f2d4034e9991
