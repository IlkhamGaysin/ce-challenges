#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int p = -1, n = 0, cp = -1, rd = 0, sbs = 32, i = 0;
	char c, d = '|';
	char *sb = malloc(sbs);

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (p == -1) {
			if (c == '\n') {
				if (cp != -1)
					p = cp;
				else
					p = rd;
				sb[p] = d;
				sb[n] = '\0';
				puts(sb);
				cp = -1;
				continue;
			} else if (c == 'C') {
				cp = n;
			} else if (c == '_') {
				rd = n;
			}
			if (n == sbs-1) {
				sbs += sbs/2;
				sb = realloc(sb, sbs);
			}
			sb[n++] = c;
		} else {
			if (c == '\n') {
				if (cp != -1) {
					if (cp < p)
						d = '/';
					else if (cp > p)
						d = '\\';
					else
						d = '|';
					p = cp;
				} else {
					if (rd < p)
						d = '/';
					else if (rd > p)
						d = '\\';
					else
						d = '|';
					p = rd;
				}
				sb[p] = d;
				puts(sb);
				i = 0;
				cp = -1;
				continue;
			} else if (p - i < 2 && p - i > -2) {
				if (c == 'C') {
					cp = i;
				} else if (c == '_') {
					rd = i;
				}
			}
			sb[i++] = c;
		}
	}
	return 0;
}
