#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF)
		printf("%d\n", 1 - (a & 1));
	return 0;
}
