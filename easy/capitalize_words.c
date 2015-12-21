#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	char c;
	bool first = true, nl = true;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		nl = c == '\n';
		if (first && islower(c)) {
			putchar(toupper(c));
			first = false;
		} else {
			putchar(c);
			first = isspace(c);
		}
	}
	if (!nl)
		putchar('\n');
	return 0;
}
