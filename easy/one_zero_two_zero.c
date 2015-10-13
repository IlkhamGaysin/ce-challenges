#include <stdbool.h>
#include <stdio.h>

bool xz(unsigned n, unsigned a) {
	while (a) {
		if (a % 2 == 0) {
			if (n)
				n--;
			else
				return false;
		}
		a >>= 1;
	}
	return n == 0;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	unsigned nz, x, i;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d %d", &nz, &x) != EOF) {
		unsigned r = 0;
		for (i = 1 << nz; i <= x; i++)
			if (xz(nz, i))
				r++;
		printf("%d\n", r);
	}
	return 0;
}
