#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c, longest[80], current[80];
	int l = 0, i = 0, j;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (c == ' ' || c == '\n') {
			if (i > l) {
				for (j = 0; j < i; j++)
					longest[j] = current[j];
				longest[i] = '\0';
				l = i;
			}
			i = 0;
			if (c == '\n') {
				puts(longest);
				l = 0;
			}
		} else {
			current[i++] = c;
		}
	}
	return 0;
}
