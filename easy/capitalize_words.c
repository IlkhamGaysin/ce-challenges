#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;
	bool first = true;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF)
		if (first && islower(c)) {
			putchar(toupper(c));
			first = false;
		} else {
			putchar(c);
			first = isspace(c);
		}
	return 0;
}
