#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, i, r[20];
	char c;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d ", &r[0]) != EOF) {
		int n = 1;
		while ((c = getc(fp)) != '|') {
			ungetc(c, fp);
			fscanf(fp, "%d ", &r[n++]);
		}
		do {
			getc(fp);
			for (i = 0; i < n; i++) {
				fscanf(fp, " %d", &a);
				if (a > r[i])
					r[i] = a;
			}
		} while (getc(fp) == ' ');
		printf("%d", r[0]);
		for (i = 1; i < n; i++) {
			printf(" %d", r[i]);
		}
		putchar('\n');
	}
	return 0;
}