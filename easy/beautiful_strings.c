#include <stdio.h>
#include <stdlib.h>

static int cmpint(const void *p1, const void *p2)
{
	return *(int*)p1 - *(int*)p2;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	char b[1024];

	fp = fopen(*++argv, "r");
	while (fgets(b, 1024, fp) != 0) {
		int c[26] = {}, p = 0, s = 0;
		while (b[p] != '\0') {
			int cc = b[p] & 223;
			p++;
			if (cc >= 65 && cc <= 90)
				c[cc - 65]++;
		}
		qsort(&c, 26, sizeof(int), cmpint);
		p = 25;
		while (c[p] > 0 && p >= 0) {
			s += (p + 1) * c[p];
			p--;
		}
		printf("%d\n", s);
	}
	return 0;
}
