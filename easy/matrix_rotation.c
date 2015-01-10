#include <stdio.h>

int sq(int a) {
	int i;
	for (i = 1; i*i < a; i++);
	return i;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	char c, m[100], t[100];
	int n = 0, l, i, j;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF || n) {
		if (c == '\n' || c == EOF) {
			l = sq(n);
			for (i = 0; i < l; i++)
				for (j = 0; j < l; j++)
					t[i*l + j] = m[j*l + i];
			for (i = 0; i < l; i++)
				for (j = l-1; j >= 0; j--) {
					putchar(t[i*l + j]);
					if (i != l-1 || j > 0)
						putchar(' ');
					else
						putchar('\n');
				}
			n = 0;
		} else if (c != ' ') {
			m[n++] = c;
		}
	}
	return 0;
}
