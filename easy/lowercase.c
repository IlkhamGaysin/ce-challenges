#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	bool p = true;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		putchar(tolower(c));
		p = c == '\n';
	}
	if (!p)
		putchar('\n');
	return 0;
}
