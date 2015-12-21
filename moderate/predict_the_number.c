#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	unsigned long long int a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%llu", &a) != EOF) {
		int i = 0;
		while (a) {
			i += a & 1;
			a >>= 1;
		}
		printf("%d\n", i % 3);
	}
	return 0;
}
