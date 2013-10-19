#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;
	int m = 1;
	char *morse = "ETIANMSURWDKGOHVF L PJBXCYZQ  54 3   2       16       7   8 90";
	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (c == '.') {
			m <<= 1;
		} else if (c == '-') {
			m = (m << 1) + 1;
		} else if (m == 1) {
			putchar(' ');
		} else {
			if (m < 64)
				putchar(morse[m-2]);
			if (c == '\n')
				putchar('\n');
			m = 1;
		}
	}
	return 0;
}
