#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int m, n;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d,%d;", &m, &n) != EOF) {
		int d, i, j;
		int *prev = calloc(n, sizeof(int));
		int *curr = calloc(n, sizeof(int));

		for (j = 0; j < n; j++)
			if (fgetc(fp) == '*')
				prev[j] = -1;
			else
				prev[j] = (j > 0 && prev[j - 1] == -1);
		for (i = 1; i < m; i++) {
			if (fgetc(fp) == '*')
				curr[0] = -1;
			else
				curr[0] = (prev[0] == -1) + (n > 1 && prev[1] == -1);
			for (j = 1; j < n; j++) {
				if (fgetc(fp) == '*')
					curr[j] = -1;
				else
					curr[j] = (curr[j - 1] == -1) + (prev[j - 1] == -1) + (prev[j] == -1) + (n > j + 1 && prev[j + 1] == -1);
				if (prev[j - 1] == -1) {
					printf("*");
				} else {
					d = prev[j - 1] + (j > 1 && curr[j - 2] == -1) + (curr[j - 1] == -1) + (curr[j] == -1) + (prev[j] == -1);
					printf("%d", d);
				}
			}
			if (prev[j - 1] == -1) {
				printf("*");
			} else {
				d = prev[j - 1] + (j > 1 && curr[j - 2] == -1) + (curr[j - 1] == -1);
				printf("%d", d);
			}
			for (j = 0; j < n; j++)
				prev[j] = curr[j];
		}
		for (j = 0; j < n; j++) {
			if (prev[j] == -1) {
				printf("*");
			} else {
				d = prev[j] + (j < n - 1 && prev[j + 1] == -1);
				printf("%d", d);
			}
		}
		printf("\n");
		free(curr);
		free(prev);
	}
	return 0;
}
