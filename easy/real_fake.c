#include <stdio.h>

const char *m[] = {"Fake", "Real"};

int main(int argc, char *argv[])
{
	FILE *fp;
	int i = 0, e = 0, o = 0;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF || i) {
		if (c == '\n' || c == EOF) {
			printf("%s\n", m[(i % 2 ? o : e) % 10 == 0]);
			i = 0;
			e = 0;
			o = 0;
		} else if (c != ' ') {
			int x = c - '0';
			e += x;
			o += x;
			if (i % 2)
				o += x;
			else
				e += x;
			i++;
		}
	}
	return 0;
}
