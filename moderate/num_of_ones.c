#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	unsigned a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF)
		printf("%d\n", __builtin_popcount(a));
	return 0;
}
