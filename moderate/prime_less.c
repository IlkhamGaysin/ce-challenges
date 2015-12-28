#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

static bool prime(int a, int *p) {
	while (*p * *p <= a) {
		if (a % *p == 0)
			return false;
		p++;
	}
	return true;
}

int main(int argc, char *argv[]) {
	FILE *fp;
	int i, n, last = 3, np = 2, ps = 8;
	int *primes = malloc(ps * sizeof(int));
	primes[0] = 2;
	primes[1] = 3;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &n) != EOF) {
		while (last + 2 < n) {
			last += 2;
			if (prime(last, primes)) {
				if (ps == np) {
					ps += ps / 2;
					primes = realloc(primes, ps * sizeof(int));
				}
				primes[np++] = last;
			}
		}
		if (n >= 2) {
			putchar('2');
			for (i = 1; i < np && primes[i] < n; i++)
				printf(",%d", primes[i]);
		}
		putchar('\n');
	}
	free(primes);
	return 0;
}
