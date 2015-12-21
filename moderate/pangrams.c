#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	char c;
	int b = 0, i;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF || b > 0) {
		if (c == '\n' || c == EOF) {
			if (b == (1 << 26) - 1) {
				puts("NULL");
			} else {
				for (i = 'a'; i <= 'z'; i++)
					if (!(b & (1 << (i - 'a'))))
						putchar(i);
				putchar('\n');
			}
			b = 0;
		} else {
			if (c >= 'a' && c <= 'z')
				b |= 1 << (c - 'a');
			else if (c >= 'A' && c <= 'Z')
				b |= 1 << (c - 'A');
		}
	}
	return 0;
}
