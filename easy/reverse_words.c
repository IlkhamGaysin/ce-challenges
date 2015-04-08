#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int sbs = 16, s = 0, rbs = 32, r = 0;
	char c;
	char *sb = malloc(sbs), *rb = malloc(rbs);

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (c == '\n')
			continue;
		while (c != '\n' && c != EOF) {
			if isblank(c) {
				while (r + s > rbs - 1) {
					rbs += rbs / 2;
					rb = realloc(rb, rbs);
				}
				while (s > 0) {
					rb[r++] = sb[--s];
				}
				rb[r++] = ' ';
				s = 0;
			} else {
				if (s == sbs - 1) {
					sbs += sbs / 2;
					sb = realloc(sb, sbs);
				}
				sb[s++] = c;
			}
			c = getc(fp);
		}
		sb[s] = '\0';
		printf("%s", sb);
		while (r > 0)
			printf("%c", rb[--r]);
		printf("\n");
		r = 0;
		s = 0;
	}
	free(rb);
	free(sb);
	return 0;
}
