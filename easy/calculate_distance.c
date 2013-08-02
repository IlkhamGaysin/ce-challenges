#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main(int argc, char *argv[])
{
    FILE *fp;
    char line[100];

    fp = fopen(*++argv, "r");
    while (fgets(line, 100, fp) != 0) {
	int a[4];
	sscanf(line, "(%d, %d) (%d, %d)", &a[0], &a[1], &a[2], &a[3]);
	int x = a[0] - a[2];
	int y = a[1] - a[3];
	int d = sqrtf((float)(x * x + y * y));
	printf("%d\n", d);
    }
    return 0;
}
