#include <math.h>
#include <stdio.h>
#include <stdlib.h>

static int cmpint(const void *p1, const void *p2) {
	return *(int *)p1 - *(int *)p2;
}

int main(int argc, char *argv[]) {
	FILE *fp;
	int i, n, f[99];

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d ", &n) != EOF) {
		int r = 0, s;
		for (i = 0; i < n; i++)
			fscanf(fp, "%d ", &f[i]);
		qsort(&f, n, sizeof(int), cmpint);
		s = f[n / 2];
		for (i = 0; i < n; i++)
			r += abs(f[i] - s);
		printf("%d\n", r);
	}
	return 0;
}
