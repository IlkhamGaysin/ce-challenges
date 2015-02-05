#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c, p = '\n';
	bool u = false;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if isalpha(c) {
			u = !u;
			if (u)
				putchar(toupper(c));
			else
				putchar(tolower(c));
		} else {
			if (c == '\n')
				u = false;
			putchar(c);
		}
		p = c;
	}
	if (p != '\n')
		putchar('\n');
	return 0;
}
