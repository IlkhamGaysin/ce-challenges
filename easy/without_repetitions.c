#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	char c, p = '\0';

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (c != p)
			putchar(c);
		p = c;
	}
	if (p != '\n')
		putchar('\n');
	return 0;
}
