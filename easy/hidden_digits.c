#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
    FILE *fp;
    char c;
    bool n = false;

    fp = fopen(*++argv, "r");
    while ((c = getc(fp)) != EOF) {
	if (c == '\n') {
	    if (!n) {
		printf("NONE");
	    }
	    printf("\n");
	    n = false;
	} else if (c >= '0' && c <= '9') {
	    printf("%c", c);
	    n = true;
	} else if (c >= 'a' && c <= 'j') {
	    printf("%c", c - 'a' + '0');
	    n = true;
	}
    }
    return 0;
}
