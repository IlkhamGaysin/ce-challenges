#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int n = 0;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF || n > 0)
		if (c == '\n' || c == EOF) {
			printf("%d\n", n);
			n = 0;
		} else {
			n += c - '0';
		}
	return 0;
}
