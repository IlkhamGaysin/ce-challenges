#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>

int powi(int a, int b)
{
	int i, ret = 1;
	for (i = 0; i < b; i++) {
		ret *= a;
	}
	return ret;
}

bool ugly(int j, int i, int *n)
{
	int k = 0, cj = j;
	int64_t s = 0, c = n[0];
	bool p = true;

	while (k < i - 1) {
		int ops = cj % 3;
		cj /= 3;
		if (ops == 0) {
			c *= 10;
		} else {
			if (p) {
				s += c;
				if (ops == 1)
					p = false;
			} else {
				s -= c;
				if (ops == 2)
					p = true;
			}
			c = 0;
		}
		k++;
		c += n[k];
	}
	if (p)
		s += c;
	else
		s -= c;
	if (s % 2 == 0 || s % 3 == 0 || s % 5 == 0 || s % 7 == 0)
		return true;
	return false;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		int i = 0, j, u = 0, n[13];
		while (c != '\n' && c != EOF) {
			n[i++] = c - 48;
			c = getc(fp);
		}
		for (j = 0; j < powi(3, i - 1); j++) {
			if (ugly(j, i, n))
				u++;
		}

		printf("%d\n", u);
	}
	return 0;
}
