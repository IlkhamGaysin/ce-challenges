#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int kernighan(int a) {
	int r = 0;
	for (; a > 0; a &= a - 1)
		r++;
	return r;
}

int numCar(int *m, int n) {
	int i, r = 0;
	for (i = 0; i < n; i++)
		r += kernighan(m[i] & ((1 << n) - 1));
	return r;
}

bool isNum(int *m, int n, int x, int y, int nc) {
	int r = 0, i;
	for (i = x; i < x + n; i++) {
		r += kernighan(m[i] & (((1 << n) - 1) << y));
		if (r > nc)
			return false;
	}
	return r == nc;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	int r[25];
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		int n = 1, i = 1, j, rm, rn;
		r[0] = c - '0';
		while ((c = getc(fp)) == '0' || c == '1') {
			n++;
			r[0] = (r[0] << 1) + c - '0';
		}
		do {
			while ((c = getc(fp)) != '0' && c != '1')
				if (c == '\n' || c == EOF)
					goto trailing;
			r[i] = c - '0';
			while ((c = getc(fp)) == '0' || c == '1')
				r[i] = (r[i] << 1) + c - '0';
			if (i < 24)
				i++;
		} while (c != '\n' && c != EOF);
trailing:
		for (rm = 1; rm <= n; rm++) {
			bool f = false;
			rn = numCar(r, rm);
			for (i = 0; i < n - rm + 1; i++) {
				for (j = 0; j < n - rm + 1; j++)
					if (!isNum(r, rm, i, j, rn)) {
						f = true;
						break;
					}
				if (f)
					break;
			}
			if (!f)
				break;
		}
		printf("%dx%d, %d\n", rm, rm, rn);
	}
	return 0;
}