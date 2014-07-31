#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		char *z = malloc(10);
		int n = 0, i;
		while isdigit(c) {
			z[n] = c-48;
			n++;
			c = getc(fp);
		}
		c = getc(fp);
		int num = 0, op = 0;
		bool sg = true;
		for (i = 0; i < n;) {
			if isalpha(c) {
				op = 10*op + z[i];
				i++;
			} else {
				if (sg)
					num += op;
				else
					num -= op;
				op = 0;
				sg = c == '+';
			}
			c = getc(fp);
		}
		if (sg)
			num += op;
		else
			num -= op;
		free(z);
		printf("%d\n", num);
	}
	return 0;
}
