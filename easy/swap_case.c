#include <ctype.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c, p = '\n';

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if islower(c)
			putchar(toupper(c));
		else if isupper(c)
			putchar(tolower(c));
		else
			putchar(c);
		p = c;
	}
	if (p != '\n')
		putchar('\n');
	return 0;
}
