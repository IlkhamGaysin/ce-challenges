#include <stdio.h>

void skip(int a, FILE *fp) {
	int i;
	for (i = 0; i < a; i++)
		fgetc(fp);
}

int main(int argc, char *argv[])
{
	FILE *fp;
	int h;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d ", &h) != EOF) {
		char c;
		int l = 0;
		while (1) {
			int cr = (l+h) % 2;
			fscanf(fp, "%c", &c);
			if (c == 'L') {
				h = (l+h)/2 + cr - 1;
				skip(5, fp);
			} else if (c == 'H') {
				l = (l+h)/2 + cr + 1;
				skip(6, fp);
			} else {
				printf("%d\n", (l+h)/2 + cr);
				skip(4, fp);
				break;
			}
		}
	}
	return 0;
}
