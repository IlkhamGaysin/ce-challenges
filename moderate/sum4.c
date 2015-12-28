#include <stdio.h>
#include <stdlib.h>

static int numzero(int *ib, int i, int c, int z) {
	if (c == 0) {
		if (z == 0)
			return 1;
		else
			return 0;
	}
	if (i < c)
		return 0;
	return numzero(ib + 1, i - 1, c - 1, z + ib[0]) + numzero(ib + 1, i - 1, c, z);
}

int main(int argc, char *argv[]) {
	FILE *fp;
	int a, i = 0, ibs = 32;
	int *ib = malloc(ibs * sizeof(int));
	char c;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		if (i == ibs) {
			ibs += ibs / 2;
			ib = realloc(ib, ibs * sizeof(int));
		}
		ib[i++] = a;
		if ((c = getc(fp)) == '\n' || c == EOF) {
			printf("%d\n", numzero(ib, i, 4, 0));
			i = 0;
		}
	}
	free(ib);
	return 0;
}
