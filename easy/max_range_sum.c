#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int n, d[99];

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d;", &n) != EOF) {
		int m = 0, cm = 0, c, i = 0;
		do {
			fscanf(fp, "%d", &c);
			if (i < n) {
				cm += c;
				d[i++] = c;
			} else {
				cm += c - d[i % n];
				d[i++ % n] = c;
			}
			if (i >= n && cm > m)
				m = cm;
		} while (getc(fp) == ' ');
		printf("%d\n", m);
	}
	return 0;
}
