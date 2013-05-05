#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int d;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%x", &d) != EOF)
		printf("%d\n", d);
	return 0;
}
