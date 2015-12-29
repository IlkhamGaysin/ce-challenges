#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	unsigned ix = 0, ibs = 16, num = 0;
	char c;
	bool swi = false;
	unsigned *ib = malloc(ibs * sizeof(unsigned));

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		switch (c) {
		case '0':
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
			num = 10 * num + ((unsigned)c - '0');
			break;
		case ' ':
			if (ix == ibs) {
				ibs += ibs / 2;
				ib = realloc(ib, ibs * sizeof(unsigned));
			}
			ib[ix++] = num;
			num = 0;
			break;
		case ':':
			swi = true;
			break;
		default:
			num = 0;
			ix = 0;
		}
		if (swi) {
			unsigned a, b, i, temp;
			do {
				fscanf(fp, " %d-%d", &a, &b);
				temp = ib[a];
				ib[a] = ib[b];
				ib[b] = temp;
			} while ((c = getc(fp)) == ',');
			printf("%d", ib[0]);
			for (i = 1; i < ix; i++) {
				printf(" %d", ib[i]);
			}
			printf("\n");
			ix = 0;
			swi = false;
		}
	}
	free(ib);
	return 0;
}
