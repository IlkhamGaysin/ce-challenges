#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	bool w = true;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		while (!((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) && c != '\n' && c != EOF)
			c = getc(fp);
		while (c != '\n' && c != EOF) {
			if (c >= 'A' && c <= 'Z')
				c |= 32;
			if (c >= 'a' && c <= 'z') {
				if (!w) {
					putchar(' ');
					w = true;
				}
				putchar(c);
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
