#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, b, c;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d,%d,%d", &a, &b, &c) != EOF) {
		if (!!(a & 1 << --b) == !!(a & 1 << --c))
			puts("true");
		else
			puts("false");
	}
	return 0;
}
