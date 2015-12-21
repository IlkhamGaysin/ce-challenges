#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int sbs = 16, s = 0;
	char c;
	char *sb = malloc(sbs);

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if isdigit(c) {
			sb[s++] = c;
			while isalpha(c = getc(fp)) {
				if (s == sbs - 2) {
					sbs += sbs / 2;
					sb = realloc(sb, sbs);
				}
				sb[s++] = c;
			}
			if isdigit(c) {
				sb[s] = sb[0];
				sb[0] = c;
			} else {
				sb[s] = c;
			}
			sb[s + 1] = '\0';
			printf("%s", sb);
		} else {
			putchar(c);
		}
		s = 0;
	}
	if (c != '\n')
		putchar('\n');
	free(sb);
	return 0;
}
