#include "vm.h"
#include "stack.h"

void vm_collect_garbage(vm_t *vm) {
        // ?
        mark(vm);
        trace(vm);
        sweep(vm);
}

void sweep(vm_t *vm) {
        // ?
        for (int i = 0; i < vm->objects->count; i++) {
                snek_object_t *obj = vm->objects->data[i];
                if (obj->is_marked) {
                        obj->is_marked = false;
                        continue;
                }
                snek_object_free(obj);
                vm->objects->data[i] = NULL;
        }
        stack_remove_nulls(vm->objects);
}

void trace(vm_t *vm) {
        // ?
        stack_t *gray_objects = stack_new(8);
        if (!gray_objects) {
                return;
        }
        for (size_t i = 0; i < vm->objects->count; i++) {
                snek_object_t *obj = vm->objects->data[i];
                if (obj->is_marked) {
                        stack_push(gray_objects, obj);
                }
        }
        while (gray_objects->count > 0) {
                trace_blacken_object(gray_objects, stack_pop(gray_objects));
        }
        stack_free(gray_objects);
}

void trace_blacken_object(stack_t *gray_objects, snek_object_t *obj) {
        // ?
        if (obj->kind == INTEGER || obj->kind == FLOAT || obj->kind == STRING) {
                return;
        }
        if (obj->kind == VECTOR3) {
                trace_mark_object(gray_objects, obj->data.v_vector3.x);
                trace_mark_object(gray_objects, obj->data.v_vector3.y);
                trace_mark_object(gray_objects, obj->data.v_vector3.z);
        }
        if (obj->kind == ARRAY) {
                for (size_t i = 0; i < obj->data.v_array.size; i++) {
                        trace_mark_object(gray_objects,
                                          obj->data.v_array.elements[i]);
                }
        }
}

void mark(vm_t *vm) {
        // ?
        for (size_t i = 0; i < vm->frames->count; i++) {
                // data is a void **
                // array of pointer to void
                // to "catch" a frame_t, we create a temporairy var
                frame_t *frame = vm->frames->data[i];
                for (size_t j = 0; j < frame->references->count; j++) {
                        snek_object_t *obj = frame->references->data[j];
                        obj->is_marked = true;
                }
        }
}

void vm_free(vm_t *vm) {
        for (int i = 0; i < vm->frames->count; i++) {
                frame_free(vm->frames->data[i]);
        }
        stack_free(vm->frames);
        for (int i = 0; i < vm->objects->count; i++) {
                snek_object_free(vm->objects->data[i]);
        }
        stack_free(vm->objects);
        free(vm);
}

// don't touch below this line

vm_t *vm_new() {
        vm_t *vm = malloc(sizeof(vm_t));
        if (vm == NULL) {
                return NULL;
        }

        vm->frames = stack_new(8);
        vm->objects = stack_new(8);
        return vm;
}

void vm_track_object(vm_t *vm, snek_object_t *obj) {
        stack_push(vm->objects, obj);
}

void vm_frame_push(vm_t *vm, frame_t *frame) { stack_push(vm->frames, frame); }

frame_t *vm_new_frame(vm_t *vm) {
        frame_t *frame = malloc(sizeof(frame_t));
        frame->references = stack_new(8);

        vm_frame_push(vm, frame);
        return frame;
}

void frame_free(frame_t *frame) {
        stack_free(frame->references);
        free(frame);
}
