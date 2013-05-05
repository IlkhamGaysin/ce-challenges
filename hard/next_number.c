#include <stdio.h>
#include <stdint.h>
#include <inttypes.h>

int64_t dsig(int64_t a)
{
	int i, d[9] = {0};
	int64_t b = a;
	while (b) {
		int r = b%10;
		if (r)
			d[r-1] += 1;
		b /= 10;
	}
	for (i = 0; i < 9; i++) {
		b *= 8;
		b += d[i];
	}
	return b;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	int64_t a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%" PRId64, &a) != EOF) {
		int64_t b = a + 1;
		while (dsig(a) != dsig(b))
			b++;
		printf("%" PRId64 "\n", b);
	}
	return 0;
}
