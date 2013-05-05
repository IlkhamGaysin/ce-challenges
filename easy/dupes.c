#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, b = -1;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		if (a != b) {
			if (b >= 0)
				printf(",");
			printf("%d", a);
			b = a;
		}
		if (getc(fp) != ',') {
			printf("\n");
			b = -1;
		}
	}
	return 0;
}
