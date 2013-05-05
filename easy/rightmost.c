#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c[256];

	fp = fopen(*++argv, "r");
	while ((c[0] = getc(fp)) != EOF) {
		char d;
		int i = 0, j = -1;
		while (c[i] != ',') {
			c[++i] = getc(fp);
		}
		d = getc(fp);
		while (--i)
			if (c[i] == d)
				j = i;
		printf("%d\n", j);
		getc(fp);
	}
	return 0;
}
