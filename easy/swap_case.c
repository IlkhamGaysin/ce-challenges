#include <ctype.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF)
		if islower(c)
			putchar(toupper(c));
		else if isupper(c)
			putchar(tolower(c));
		else
			putchar(c);
	return 0;
}
