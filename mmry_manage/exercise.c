#include "exercise.h"
#include <stdbool.h>
#include <stddef.h>

snek_object_t *snek_add(snek_object_t *a, snek_object_t *b) {
        // ?
        if (!a || !b) {
                return NULL;
        }

        if (a->kind == INTEGER) {
                if (b->kind == INTEGER) {
                        // return new_snek_integer(a->data.v_int +
                        // b->data.v_int);
                        return new_snek_integer(a->data.v_int);
                }
                if (b->kind == FLOAT) {
                        return new_snek_float((float)a->data.v_int +
                                              b->data.v_float);
                }
                return NULL;
        }

        if (a->kind == FLOAT) {
                if (b->kind == INTEGER) {
                        return new_snek_float(a->data.v_float +
                                              (float)b->data.v_int);
                }
                if (b->kind == FLOAT) {
                        return new_snek_float(a->data.v_float +
                                              b->data.v_float);
                }
                return NULL;
        }

        if (a->kind == STRING) {
                if (b->kind != STRING) {
                        return NULL;
                } else {
                        size_t len =
                            strlen(a->data.v_string) + strlen(b->data.v_string);
                        char *tmp = calloc(len + 1, sizeof(char));
                        tmp = strcat(a->data.v_string, b->data.v_string);
                        snek_object_t *str = new_snek_string(tmp);
                        free(tmp);
                        return str;
                }
        }

        if (a->kind == VECTOR3) {
                if (b->kind != VECTOR3) {
                        return NULL;
                } else {
                        snek_object_t *new_snek_vector3(
                            snek_add(a->data.v_vector3.x + b->data.v_vector3.x),
                            snek_add(a->data.v_vector3.y + b->data.v_vector3.y),
                            snek_add(a->data.v_vector3.z +
                                     b->data.v_vector3.z));
                }
        }

        if (a->kind == ARRAY) {
                if (b->kind != ARRAY) {
                        return NULL;
                } else {
                        snek_object_t *arr =
                            new_snek_array(snek_length(a) + snek_length(b));
                        for (size_t i = 0; i < a->data.v_array.size; i++) {
                                if (!snek_array_set(arr, i, a)) {
                                        arr = snek_array_get(a, i);
                                }
                        }
                        for (size_t i = 0; i < b->data.v_array.size; i++) {
                                if (!snek_array_set(arr, i, b)) {
                                        arr = snek_array_get(b, i);
                                }
                        }
                }
                return arr;
        }

        return NULL;
}
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
