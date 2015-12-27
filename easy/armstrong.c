#include <stdio.h>

static int powi(int a, int b) {
	int i, ret = 1;
	for (i = 0; i < b; i++)
		ret *= a;
	return ret;
}

int main(int argc, char *argv[]) {
	FILE *fp;
	int n;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d;", &n) != EOF) {
		int a, e = 0, s = 0;
		for (a = n; a; a /= 10)
			e++;
		for (a = n; a; a /= 10)
			s += powi(a % 10, e);
		puts(n == s ? "True" : "False");
	}
	return 0;
}
