#include <ctype.h>
#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char b[256];

	fp = fopen(*++argv, "r");
	while (fgets(b, 31, fp) != 0) {
		int p = strlen(b) - 1, q = p;
		while (p >= -1) {
			if (p == -1 || isblank(b[p])) {
				int r = p;
				while (++r < q)
					putchar(b[r]);
				if (p >= 0)
					printf(" ");
				q = p;
			}
			p--;
		}
		printf("\n");
	}
	return 0;
}
