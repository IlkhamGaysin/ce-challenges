#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int n, m, ls = 100, i, j, c;
	char *l = malloc(ls);

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d %d\n", &n, &m) != EOF) {
		if (m == 0) {
			printf("%d\n", n);
		} else if (m == 1) {
			printf("%d\n", n - 1);
		} else {
			if (n > ls) {
				ls = n;
				l = realloc(l, ls);
			}
			for (i = 0; i < n; i++)
				l[i] = 0;
			for (i = 0; i <= m % 2; i++) {
				for (j = 1; j < n; j += 2)
					l[j] = 1;
				for (j = 2; j < n; j += 3)
					if (l[j] != 0)
						l[j] = 0;
					else
						l[j] = 1;
			}
			if (l[n - 1] != 0)
				l[n - 1] = 0;
			else
				l[n - 1] = 1;
			c = n;
			for (i = 0; i < n; i++)
				c -= l[i];
			printf("%d\n", c);
		}
	}
	free(l);
	return 0;
}
