#include <stdbool.h>
#include <stdio.h>

bool pali(int a) {
	int l[16], c = a, x = 0, y = 0;
	while (c) {
		l[y++] = c % 10;
		c /= 10;
	}
	while (true) {
		if (l[x] != l[y - 1])
			return false;
		if (y - x < 3)
			return true;
		x++;
		y--;
	}
}

int main(int argc, char *argv[]) {
	FILE *fp;
	int l, r;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d %d", &l, &r) != EOF) {
		int i, n = 0;
		for (i = l; i <= r; i++) {
			int j, prev = -1;
			for (j = i; j <= r; j++) {
				int k, p;
				if (prev > -1) {
					p = prev;
					if (pali(j))
						p++;
				} else {
					p = 0;
					for (k = i; k <= j; k++) {
						if (pali(k))
							p++;
					}
				}
				if (p % 2 == 0)
					n++;
				prev = p;
			}
		}
		printf("%d\n", n);
	}
	return 0;
}
