#include <stdio.h>

int main(int argc, char *argv[]) {
	FILE *fp;
	int s = 0, a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF)
		s += a;
	printf("%d\n", s);
	return 0;
}
