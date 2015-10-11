#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char line[72];
	unsigned i, n, r;

	fp = fopen(*++argv, "r");
	while (fgets(line, 72, fp) != 0) {
		r = 1;
		n = strlen(line) / 2;
		for (i = 0; i < n; i++) {
			if ((line[i] == 'A' && line[i+n] == 'B') ||
			    (line[i] == 'B' && line[i+n] == 'A')) {
				r = 0;
				break;
			} else if (line[i] == '*' && line[i+n] == '*') {
				r *= 2;
			}
		}
		printf("%d\n", r);
	}
	return 0;
}