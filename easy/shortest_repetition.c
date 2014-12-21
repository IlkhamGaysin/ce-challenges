#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int p = 1, i = 0, sbs = 32;
	char c;
	char *sb = malloc(sbs);

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (c == '\n') {
			printf("%d\n", p);
			p = 1;
			i = 0;
			continue;
		}
		if (i == sbs) {
			sbs += sbs/2;
			sb = realloc(sb, sbs);
		}
		if (i > 0 && sb[i-p] != c)
			p = i+1;
		sb[i++] = c;
	}
	free(sb);
	return 0;
}
