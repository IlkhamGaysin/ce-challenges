#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;
	bool u = false;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF)
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
	return 0;
}
