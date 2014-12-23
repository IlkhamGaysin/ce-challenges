#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct item {
	int		num;
	int		pos;
};

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, i = 0, j = 1, ibs = 32;
	struct item *ib = malloc(ibs * sizeof(struct item));

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		bool f = false;
		int k, m = a;
		ib[i].num = a;
		ib[i++].pos = j++;
		while (getc(fp) == ' ') {
			fscanf(fp, "%d", &a);
			for (k = 0; k < i; k++) {
				if (ib[k].num == a) {
					ib[k].pos = 0;
					break;
				}
			}
			if (k == i) {
				if (m < a)
					m = a;
				if (i == ibs) {
					ibs += ibs/2;
					ib = realloc(ib, ibs * sizeof(struct item));
				}
				ib[i].num = a;
				ib[i++].pos = j;
			}
			j++;
		}
		for (k = 0; k < i; k++) {
			if (ib[k].pos != 0 && ib[k].num <= m) {
				m = ib[k].num;
				j = ib[k].pos;
				f = true;
			}
		}
		if (f) {
			printf("%d\n", j);
		} else {
			puts("0");
		}
		i = 0;
		j = 1;
	}
	free(ib);
	return 0;
}
