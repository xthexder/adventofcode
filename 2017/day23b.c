
#include <stdio.h>

int main() {

    int b = 109900;
    int h = 0;

    while (1) {
        int f = 1;
        int d = 2;
        do {
            int e = 2;
            // do {
            //     if (d * e == b) {
            //         f = 0;
            //     }
            //     e++;
            // } while (e != b);
            if (b % d == 0) {
                f = 0;
            }

            d++;
        } while (d != b);

        if (f == 0) {
            h++;
        }
        if (b == 126900) {
            break;
        }
        b += 17;
    }
    printf("h=%d\n", h);
}
