#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, n, i = 0, ibs = 16;
	int *ib = malloc(ibs * sizeof(int));

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		if (i == ibs) {
			ibs += ibs / 2;
			ib = realloc(ib, ibs * sizeof(int));
		}
		ib[i++] = a;
		if (getc(fp) == ';') {
			int j = 0, k = 0, l = 0;
			fscanf(fp, "%d", &n);
			while (j + n <= i) {
				for (k = j + n - 1; k >= j; k--) {
					printf("%d", ib[k]);
					if (++l < i)
						printf(",");
				}
				j += n;
			}
			for (k = j; k < i; k++) {
				printf("%d", ib[k]);
				if (++l < i)
					printf(",");
			}
			printf("\n");
			getc(fp);
			i = 0;
		}
	}
	free(ib);
	return 0;
}
