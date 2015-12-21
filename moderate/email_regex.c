#include <regex.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	regex_t email_regex;
	int sbs = 48, n = 0;
	char c;
	char *sb = malloc(sbs);

	regcomp(&email_regex, "^((\"[0-9A-Za-z@.+-=]+\")|([0-9A-Za-z.+-=]+))"
			      "@\\w*(\\w+\\.)+\\w{2,4}$",
		REG_EXTENDED | REG_NOSUB);

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF || n > 0) {
		if (c == '\n' || c == EOF) {
			sb[n] = '\0';
			if (!regexec(&email_regex, sb, 0, NULL, 0)) {
				puts("true");
			} else {
				puts("false");
			}
			n = 0;
		} else {
			if (n == sbs - 1) {
				sbs += sbs / 2;
				sb = realloc(sb, sbs);
			}
			sb[n++] = c;
		}
	}
	free(sb);
	return 0;
}
