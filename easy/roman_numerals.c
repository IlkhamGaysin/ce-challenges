#include <stdio.h>

const int ronum[13] = { 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1 };
const char *rostr[13] = { "M", "CM", "D", "CD", "C", "XC",
			  "L", "XL", "X", "IX", "V", "IV", "I" };

int main(int argc, char *argv[]) {
	FILE *fp;
	int a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		int i = 0;
		while (a > 0) {
			if (a >= ronum[i]) {
				printf("%s", rostr[i]);
				a -= ronum[i];
			} else {
				i++;
			}
		}
		printf("\n");
	}
	return 0;
}
