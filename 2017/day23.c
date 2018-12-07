
#include <stdio.h>

int main() {
    int mulCount = 0;

    int a = 0;
    int b = 0;
    int c = 0;
    int d = 0;
    int e = 0;
    int f = 0;
    int g = 0;
    int h = 0;

    line0:	    b = 99;
    line1:	    c = b;
    line2:	    if (a != 0) goto line4;
    line3:	    if (1 != 0) goto line8;
    line4:	    b *= 100; mulCount++;
    line5:	    b -= -100000;
    line6:	    c = b;
    line7:	    c -= -17000;
    line8:	    f = 1;
    line9:	    d = 2;
    line10:	    e = 2;
    line11:	    g = d;
    line12:	    g *= e; mulCount++;
    line13:	    g -= b;
    line14:	    if (g != 0) goto line16;
    line15:	    f = 0;
    line16:	    e -= -1;
    line17:	    g = e;
    line18:	    g -= b;
    line19:	    if (g != 0) goto line11;
    line20:	    d -= -1;
    line21:	    g = d;
    line22:	    g -= b;
    line23:	    if (g != 0) goto line10;
    line24:	    if (f != 0) goto line26;
    line25:	    h -= -1;
    line26:	    g = b;
    line27:	    g -= c;
    line28:	    if (g != 0) goto line30;
    line29:	    if (1 != 0) goto line32;
    line30:	    b -= -17;
    line31:	    if (1 != 0) goto line8;
    line32:
    printf("Registers: a=%d b=%d c=%d d=%d e=%d f=%d g=%d h=%d\n", a, b, c, d, e, f, g, h);
    printf("mulCount = %d\n", mulCount);
}
