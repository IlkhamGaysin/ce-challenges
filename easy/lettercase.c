#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;
	float count = 0, lower = 0, ratio;

	fp = fopen(*++argv, "r");
	while (1) {
		if (fscanf(fp, "%c", &c) == EOF) {
			if (count == 0)
				break;
			else
				c = '\n';
		}
		if (c == '\n') {
			ratio = 100 * lower / count;
			printf("lowercase: %.2f uppercase: %.2f\n", ratio, 100-ratio);
			count = 0;
			lower = 0;
		} else if (c >= 'a' && c <= 'z') {
			lower++;
			count++;
		} else if (c >= 'A' && c <= 'Z') {
			count++;
		}
	}
	return 0;
}
