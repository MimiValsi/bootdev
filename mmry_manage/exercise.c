#include "exercise.h"
#include <iterator>
#include <stdbool.h>
#include <stddef.h>

// void snek_zero_out(void *ptr, snek_object_kind_t kind) {
//
//         // need to cast to the right struct depending the kind variable
//         switch (kind) {
//         case INTEGER:
//                 ((snek_int_t *)ptr)->value = 0;
//                 break;
//         case FLOAT:
//                 ((snek_float_t *)ptr)->value = 0.0;
//                 break;
//         case BOOL:
//                 ((snek_bool_t *)ptr)->value = 0;
//         }
// }

bool snek_array_set(snek_object_t *snek_obj, size_t index,
                    snek_object_t *value) {
        // ?
        if (!snek_obj || !value) {
                return false;
        }
        if (snek_obj->kind != ARRAY) {
                return false;
        }
        if (index > snek_obj->data.v_array.size) {
                return false;
        }
        snek_obj->data.v_array.elements[index] = value;

        return true;
}

snek_object_t *snek_array_get(snek_object_t *snek_obj, size_t index) {
        if (!snek_obj) {
                return NULL;
        }
        if (snek_obj->kind != ARRAY) {
                return NULL;
        }
        if (index >= snek_obj->data.v_array.size) {
                return NULL;
        }
        return snek_obj->data.v_array.elements[index];
}
// void format_object(snek_object_t obj, char *buffer) {
//         // ?
//         switch (obj.kind) {
//         case INTEGER:
//                 sprintf(buffer, "int:%d", obj.data.v_int);
//                 break;
//         case STRING:
//                 sprintf(buffer, "string:%s", obj.data.v_string);
//                 break;
//         }
// }
//
// // don't touch below this line'
//
// snek_object_t new_integer(int i) {
//         return (snek_object_t){.kind = INTEGER, .data = {.v_int = i}};
// }
//
// snek_object_t new_string(char *str) {
//         // NOTE: We will learn how to copy this data later.
//         return (snek_object_t){.kind = STRING, .data = {.v_string = str}};
// }
