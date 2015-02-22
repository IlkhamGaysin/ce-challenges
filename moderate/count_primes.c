#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

bool prime(int a, int *p) {
	while (*p * *p <= a) {
		if (a % *p == 0)
			return false;
		p++;
	}
	return true;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	int psize = 128, plast = 1, pmax = 3;
	int *primes = malloc(psize * sizeof(int));
	int i, x, y, c;
	primes[0] = 2;
	primes[1] = 3;
	for (i = primes[plast] + 2; plast < psize - 1; i += 2)
		if (prime(i, primes))
			primes[++plast] = i;
	pmax = i - 2;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d,%d", &x, &y) != EOF) {
		while (y > pmax) {
			psize += psize / 2;
			primes = realloc(primes, psize * sizeof(int));
			for (i = primes[plast] + 2; plast < psize - 1; i += 2)
				if (prime(i, primes))
					primes[++plast] = i;
			pmax = i - 2;
		}
		c = 0;
		for (i = 0; primes[i] <= y; i++)
			if (x <= primes[i])
				c++;
		printf("%d\n", c);
	}
	free(primes);
	return 0;
}
