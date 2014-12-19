#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int i = 0, sbs = 32;
	char c;
	char *sb = malloc(sbs);
	char *sbp = malloc(sbs);
	sbp[0] = '\0';

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (c == '\n') {
			printf("%s\n", sbp);
			sbp[0] = '\0';
			i = 0;
		} else if (c == ' ') {
			sb[i] = '\0';
			free(sbp);
			sbp = sb;
			sb = malloc(sbs);
			i = 0;
		} else {
			if (i == sbs-1) {
				sbs += sbs/2;
				sb = realloc(sb, sbs);
				sbp = realloc(sbp, sbs);
			}
			sb[i++] = c;
		}
	}
	free(sbp);
	free(sb);
	return 0;
}
