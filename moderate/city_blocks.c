#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	float st[100], av[100];
	int nst = 0, nav = 0, i, j;

	fp = fopen(*++argv, "r");
	while (getc(fp) != EOF) {
		do {
			fscanf(fp, "%f", &st[nst++]);
		} while (getc(fp) != ')');
		getc(fp);
		getc(fp);
		do {
			fscanf(fp, "%f", &av[nav++]);
		} while (getc(fp) != ')');
		getc(fp);
		if (nst < 2 || nav < 2) {
			puts("0");
		} else {
			int its = 0, stl, avl;
			if (st[0])
				for (i = nst - 1; i >= 0; i--)
					st[i] -= st[0];
			if (av[0])
				for (i = nav - 1; i >= 0; i--)
					av[i] -= av[0];
			stl = st[nst - 1];
			avl = av[nav - 1];
			for (i = 0; i < nav; i++)
				av[i] = av[i] * stl / avl;
			for (i = 0, j = 0; i < nst && j < nav;) {
				if (st[i] == av[j]) {
					its++;
					i++;
					j++;
				} else if (st[i] < av[j]) {
					i++;
				} else {
					j++;
				}
			}
			printf("%d\n", nst + nav - its - 1);
		}
		nst = 0;
		nav = 0;
	}
	return 0;
}
