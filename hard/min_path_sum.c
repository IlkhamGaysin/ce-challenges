#include <stdio.h>
#include <stdlib.h>

int min(int a, int b) {
	return (a < b) ? a : b;
}

int main(int argc, char *argv[]) {
	FILE *fp;
	int a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d%*c", &a) != EOF) {
		int i, j, d;
		int *b = malloc(sizeof(int) * a);
		fscanf(fp, "%d%*c", &d);
		b[0] = d;
		for (i = 1; i < a; i++) {
			fscanf(fp, "%d%*c", &d);
			b[i] = d + b[i - 1];
		}
		for (i = 1; i < a; i++) {
			fscanf(fp, "%d%*c", &d);
			b[0] += d;
			for (j = 1; j < a; j++) {
				fscanf(fp, "%d%*c", &d);
				b[j] = d + min(b[j], b[j - 1]);
			}
		}
		printf("%d\n", b[a - 1]);
		free(b);
	}
	return 0;
}
