#include "snekobject.h"

int main(void) {
        snek_object_t *one = new_snek_integer(1);
        snek_object_t *three = new_snek_integer(3);
        snek_object_t *four = snek_add(one, three);
}
