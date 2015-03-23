#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int i, s = 0, sbs = 16;
	char c;
	float f, g;
	char *sb = malloc(sbs);

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		while (c == '*' || c == '/' || c == '+') {
			if (s == sbs) {
				sbs += sbs / 2;
				sb = realloc(sb, sbs);
			}
			sb[s++] = c;
			getc(fp);
			c = getc(fp);
		}
		ungetc(c, fp);
		fscanf(fp, "%f ", &f);
		for (i = s - 1; i >= 0; i--) {
			fscanf(fp, "%f ", &g);
			switch (sb[i]) {
			case '*':
				f *= g;
				break;
			case '/':
				f /= g;
				break;
			default:
				f += g;
			}
		}
		printf("%d\n", (int)(f + 0.001));
		s = 0;
	}

	free(sb);
	return 0;
}
