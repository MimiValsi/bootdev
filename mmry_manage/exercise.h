<<<<<<< HEAD
#include <stdbool.h>
#include <stddef.h>

typedef struct SnekObject snek_object_t;

typedef struct SnekVector {
        snek_object_t *x;
        snek_object_t *y;
        snek_object_t *z;
} snek_vector_t;

=======
#include <stddef.h>

>>>>>>> 84524d623211a705e578a3e72876f2d4034e9991
typedef enum SnekObjectKind {
        INTEGER,
        STRING,
        FLOAT,
        BOOL,
        VECTOR3,
        ARRAY,
} snek_object_kind_t;

<<<<<<< HEAD
typedef struct {
        size_t size;
        snek_object_t **elements;
} snek_array_t;
=======
typedef struct Stack {
        size_t count;
        size_t capacity;
        void **data;
} stack_t;
>>>>>>> 84524d623211a705e578a3e72876f2d4034e9991

typedef struct SnekInt {
        char *name;
        int value;
} snek_int_t;

typedef struct SnekFloat {
        char *name;
        int value;
} snek_float_t;

typedef struct SnekBool {
        char *name;
        int value;
} snek_bool_t;

typedef union SnekObjectData {
        int v_int;
        float v_float;
        char *v_string;
        snek_vector_t v_vector3;
        snek_array_t v_array;
} snek_object_data_t;

typedef struct SnekObject {
        snek_object_kind_t kind;
        snek_object_data_t data;
} snek_object_t;

snek_object_t new_integer(int);
snek_object_t new_string(char *str);
void format_object(snek_object_t obj, char *buffer);
void snek_zero_out(void *ptr, snek_object_kind_t kind);
bool snek_array_set(snek_object_t *array, size_t index, snek_object_t *value);
snek_object_t *snek_array_get(snek_object_t *array, size_t index);
