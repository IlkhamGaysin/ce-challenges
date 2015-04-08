#include <stdio.h>
#include <stdlib.h>

int indx(char *a, char c) {
	int i = 0;
	while (a[i] != c && i < 84)
		i++;
	return i;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	int sbs = 8, s = 0, i = 0;
	char c;
	char *alpha = " !\"#$%&'()*+,-./0123456789:<=>?@"
		      "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
	char *sb = malloc(sbs);

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		while (c != ';') {
			if (s == sbs) {
				sbs += sbs / 2;
				sb = realloc(sb, sbs);
			}
			sb[s++] = c - '0';
			c = getc(fp);
		}
		while ((c = getc(fp)) != '\n' && c != EOF)
			printf("%c", alpha[(indx(alpha, c) - sb[i++ % s] + 84) % 84]);
		printf("\n");
		s = 0;
		i = 0;
	}
	free(sb);
	return 0;
}
