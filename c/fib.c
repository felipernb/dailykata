#include <stdio.h>
//O(n) in time, O(1) in space 
int fib(int n) {
    int a, b, i, temp;
    a = 0;
    b = 1;
    for (i = 0; i < n; i++) {
        temp = a;
        a = a + b;
        b = temp;
    }
    return a;
}

int main() {
    int i;
    for (i = 0; i < 10; i++) {
        printf("fib(%d) = %d\n", i, fib(i));
    }
}
