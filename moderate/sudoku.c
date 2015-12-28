#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

static int sq(int a) {
	int i;
	for (i = 1; i * i < a; i++) ;
	return i;
}

int main(int argc, char *argv[]) {
	FILE *fp;
	int a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d%*c", &a) != EOF) {
		bool valid = true;
		int d, i, j, crow;
		int *csqu = calloc(sq(a), sizeof(int));
		int *col = calloc(a, sizeof(int));

		for (i = 0; i < a; i++) {
			crow = 0;
			for (j = 0; j < a; j++) {
				fscanf(fp, "%d%*c", &d);
				d = 1 << (d - 1);
				crow |= d;
				csqu[j / sq(a)] |= d;
				col[j] |= d;
			}
			d = (1 << a) - 1;
			if (crow != d) {
				valid = false;
			}
			if ((i + 1) % sq(a) == 0) {
				for (j = 0; j < sq(a); j++) {
					if (csqu[j] != d) {
						valid = false;
					} else {
						csqu[j] = 0;
					}
				}
			}
		}
		for (i = 0; valid && i < a; i++) {
			if (col[i] != (1 << a) - 1) {
				valid = false;
			}
		}
		puts (valid ? "True" : "False");
		free(col);
		free(csqu);
	}
	return 0;
}
