#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int i = 0, sbs = 32;
	char c;
	char *sb = malloc(sbs);

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (c == ',') {
			int j;
			getc(fp);
			while ((c = getc(fp)) != '\n')
				for (j = 0; j < i; j++)
					if (sb[j] == c)
						sb[j] = ',';
			for (j = 0; j < i; j++)
				if (sb[j] != ',')
					printf("%c", sb[j]);
			printf("\n");
			i = 0;
		} else {
			if (i == sbs-1) {
				sbs += sbs/2;
				sb = realloc(sb, sbs);
			}
			sb[i++] = c;
		}
	}
	free(sb);
	return 0;
}
