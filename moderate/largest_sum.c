#include <stdio.h>
#include <limits.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, l = 0, m = INT_MIN;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		char c = getc(fp);

		if (a > m)
			m = a;
		if (a + l > m)
			m = a + l;
		if (a + l > 0)
			l += a;

		if (c == '\n') {
			printf("%d\n", m);
			l = 0;
			m = INT_MIN;
		}
	}
	return 0;
}
