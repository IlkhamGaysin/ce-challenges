#include <stdio.h>

int powi(int a, int b)
{
	int i, ret = 1;
	for (i = 0; i < b; i++)
		ret *= a;
	return ret;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	int n;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d;", &n) != EOF) {
		int a, e = 0, s = 0, r;
		for (a = n; a; a /= 10)
			e++;
		for (a = n; a; a /= 10) {
			r = a % 10;
			s += powi(r, e);
		}
		if (n == s)
			printf("True\n");
		else
			printf("False\n");
	}
	return 0;
}
