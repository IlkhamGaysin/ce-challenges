#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int s = 0, t = 0, sbs = 32, i, j;
	char c;
	char *sb = malloc(sbs), *tb = malloc(sbs);

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (c == ',') {
			bool rot = true;
			while ((c = getc(fp)) != '\n' && c != EOF) {
				if (t > s)
					rot = false;
				if (rot)
					tb[t++] = c;
			}
			if (t != s)
				rot = false;
			if (rot) {
				for (i = 0; i < s; i++) {
					rot = true;
					for (j = 0; j < s; j++) {
						if (sb[j] != tb[(i+j)%s]) {
							rot = false;
							break;
						}
					}
					if (rot)
						break;
				}
			}
			if (rot)
				puts("True");
			else
				puts("False");
			s = 0;
			t = 0;
		} else {
			if (s == sbs) {
				sbs += sbs/2;
				sb = realloc(sb, sbs);
				tb = realloc(tb, sbs);
			}
			sb[s++] = c;
		}
	}
	free(sb);
	free(tb);
	return 0;
}
