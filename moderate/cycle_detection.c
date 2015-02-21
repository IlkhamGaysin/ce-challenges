#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

bool uniq(int *ib, int s, int i) {
	int j;
	for (j = s + 1; j < i; j++)
		if (ib[s] == ib[j])
			return false;
	return true;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, s = 0, i = 0, ibs = 32;
	int *ib = malloc(ibs * sizeof(int));
	char c;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		if (i == ibs) {
			ibs += ibs / 2;
			ib = realloc(ib, ibs * sizeof(int));
		}
		ib[i++] = a;
		if ((c = getc(fp)) == '\n' || c == EOF) {
			while (i - s > 1 && uniq(ib, s, i))
				s++;
			if (s < i - 1) {
				printf("%d", ib[s]);
				for (i = s + 1; ib[i] != ib[s]; i++)
					printf(" %d", ib[i]);
			}
			printf("\n");
			s = 0;
			i = 0;
		}
	}
	free(ib);
	return 0;
}
