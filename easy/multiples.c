#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char line[48];

	fp = fopen(*++argv, "r");
	while (fgets(line, 48, fp) != 0) {
		int a, b, i = 0;
		sscanf(line, "%d,%d", &a, &b);
		while (b * ++i < a);
		printf("%d\n", b * i);
	}
	return 0;
}
