#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int a, c, i = 0, ibs = 32, j, k;
	int *ib = malloc(ibs * sizeof(int));

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d ", &a) != EOF) {
		ib[i++] = a;
		while ((c = fscanf(fp, "%d ", &a)) > 0) {
			if (i == ibs) {
				ibs += ibs / 2;
				ib = realloc(ib, ibs * sizeof(int));
			}
			ib[i++] = a;
		}
		fscanf(fp, "| %d", &a);
		for (j = 0; j < i - 1 && j < a; j++) {
			for (k = 1; k < i; k++) {
				if (ib[k - 1] > ib[k]) {
					int x = ib[k - 1];
					ib[k - 1] = ib[k];
					ib[k] = x;
				}
			}
		}
		for (j = 0; j < i; j++) {
			if (j > 0)
				putchar(' ');
			printf("%d", ib[j]);
		}
		putchar('\n');
		i = 0;
	}
	free(ib);
	return 0;
}
