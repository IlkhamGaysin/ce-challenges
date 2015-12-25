#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int n;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &n) != EOF) {
		int xo, i;
		bool jolly = true, done = false;
		int *diff = calloc(n, sizeof(int));
		if (diff == NULL) {
			puts("Out of memory");
			exit(-1);
		}
		if (n == 1)
			done = true;
		fscanf(fp, "%d", &xo);
		for (i = 1; i < n; i++) {
			int xc, x;
			fscanf(fp, "%d", &xc);
			if (done)
				continue;
			x = abs(xc - xo);
			if (x >= n || x == 0 || diff[x - 1]) {
				jolly = false;
				done = true;
			} else {
				diff[x - 1] = 1;
				xo = xc;
			}
		}
		puts(jolly ? "Jolly" : "Not jolly");
		free(diff);
	}
	return 0;
}
