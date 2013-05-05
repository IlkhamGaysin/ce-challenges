#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		int i = 0;
		while (a) {
			i += a & 1;
			a >>= 1;
		}
		printf("%d\n", i);
	}
	return 0;
}
