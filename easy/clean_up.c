#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	bool w = true;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		while (!isalpha(c) && c != '\n' && c != EOF)
			c = getc(fp);
		while (c != '\n' && c != EOF) {
			if isalpha(c) {
				if (!w) {
					putchar(' ');
					w = true;
				}
				putchar(tolower(c));
			} else if (w) {
				w = false;
			}
			c = getc(fp);
		}
		putchar('\n');
		w = true;
	}
	return 0;
}
