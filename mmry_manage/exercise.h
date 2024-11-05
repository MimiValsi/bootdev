typedef enum SnekObjectKind {
        INTEGER,
        // STRING,
        FLOAT,
        BOOL,
} snek_object_kind_t;

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
        char *v_string;
} snek_object_data_t;

typedef struct SnekObject {
        snek_object_kind_t kind;
        snek_object_data_t data;
} snek_object_t;

snek_object_t new_integer(int);
snek_object_t new_string(char *str);
void format_object(snek_object_t obj, char *buffer);
void snek_zero_out(void *ptr, snek_object_kind_t kind);
