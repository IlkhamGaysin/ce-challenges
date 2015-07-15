#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c, s[20];
	int i = 0;
	bool p = true, w = true;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (w) {
			if (c == ' ') {
				i = 0;
				w = false;
			} else {
				s[i++] = c;
			}
		} else {
			if (c == '1') {
				putchar(toupper(s[i++]));
			} else if (c == '0') {
				putchar(s[i++]);
			} else {
				putchar('\n');
				i = 0;
				w = true;
			}
			p = c == '\n';
		}
	}
	if (!p)
		putchar('\n');
	return 0;
}