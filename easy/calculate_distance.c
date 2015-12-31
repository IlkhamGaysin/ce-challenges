#include <math.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int a[4];
	char line[27];

	if (argc != 2) {
		printf("Usage: %s [FILE]\n", argv[0]);
		return 1;
	}
	fp = fopen(*++argv, "r");
	while (fgets(line, 27, fp) != 0) {
		sscanf(line, "(%d, %d) (%d, %d)", &a[0], &a[1], &a[2], &a[3]);
		int x = a[0] - a[2];
		int y = a[1] - a[3];
		int d = sqrtf((float)(x * x + y * y));
		printf("%d\n", d);
	}
	return 0;
}
