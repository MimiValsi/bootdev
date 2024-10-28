# Memory Management

* A struct stores 8 integers in its ordered fields: A through H.\
An array stores 10 of these structs, what will be the offset in bytes\
from the start of the array to the 5th element's C field?

struct Foo {
 int A; 4 bytes
 int B;
 int C;
 int D;
 int E;
 int F;
 int G;
 int H;
};
sizeof(Foo) = 8 * 4 = 32 bytes

Foo[10];
(index as 4!)
5th element = 32 * 4 = 128 bytes

offset from start
0 1 2 3 4 5
        ^
        5th element

Foo[4] = {A, B, C, D, E, F, G, H};
          0, 1, 2, 3, 4, 5, 6, 7
                ^
                3th element

Therefore, 2 * 4 = 8 bytes
Add 128 + 8 = 136 bytes!
