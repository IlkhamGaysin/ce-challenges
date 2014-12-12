#include <stdbool.h>
#include <stdio.h>

const char *phrases[] = {", yeah!", ", this is crazy, I tell ya.",
			 ", can U believe this?", ", eh?", ", aw yea.",
			 ", yo.", "? No way!", ". Awesome!"};

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;
	bool l = false;
	int i = 0;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (c == '.' || c == '!' || c == '?') {
			if (l) {
				printf("%s", phrases[i]);
				l = false;
				i = (i+1) % 8;
			} else {
				printf("%c", c);
				l = true;
			}
		} else {
			printf("%c", c);
		}
	}
	return 0;
}
