#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, i;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		for (i = 0; a > 0; a &= a - 1)
			i++;
		printf("%d\n", i);
	}
	return 0;
}
