#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	bool p = true;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if islower(c)
			putchar(toupper(c));
		else if isupper(c)
			putchar(tolower(c));
		else
			putchar(c);
		p = c == '\n';
	}
	if (!p)
		putchar('\n');
	return 0;
}
