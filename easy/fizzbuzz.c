#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int fizz, buzz, n;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d %d %d", &fizz, &buzz, &n) != EOF) {
		int i;
		for (i = 1; i <= n; i++) {
			if (i > 1)
				printf(" ");
			if (i % fizz && i % buzz) {
				printf("%d", i);
			} else {
				if (i % fizz == 0)
					printf("F");
				if (i % buzz == 0)
					printf("B");
			}
		}
		printf("\n");
	}
	return 0;
}
