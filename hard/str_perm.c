#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int next_permu(char *c, int strl)
{
	char t;
	int i, k = -1, l = 0;

	for (i = strl-2; i >= 0; i--) {
		if (c[i] < c[i+1]) {
			k = i;
			break;
		}
	}
	if (k == -1)
		return 0;
	for (i = strl-1; i > 0; i--) {
		if (c[i] > c[k]) {
			l = i;
			break;
		}
	}

	t = c[k];
	c[k] = c[l];
	c[l] = t;

	l = strl - k - 1;
	for (i = 0; (l-1)-i > i; i++) {
		t = c[strl-1-i];
		c[strl-1-i] = c[k+1+i];
		c[k+1+i] = t;
	}
	return l;
}

int ccmp(const void *b1, const void *b2)
{
	char *c1 = (char *)b1, *c2 = (char *)b2;
	if (*c1 < *c2)
		return -1;
	return *c1 > *c2;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	char line[32];

	fp = fopen(*++argv, "r");
	while (fgets(line, 31, fp) != 0) {
		int strl = strlen(line)-1;
		char *c, *w =  malloc(strl*sizeof(char));
		c = strncpy(w, line, strl);
		qsort(c, strl, sizeof(char), ccmp);
		printf("%s", c);
		while (next_permu(c, strl))
			printf(",%s", c);
		printf("\n");
		free(w);
	}
	return 0;
}
