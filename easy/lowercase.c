#include <ctype.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF)
		putchar(tolower(c));
	return 0;
}
