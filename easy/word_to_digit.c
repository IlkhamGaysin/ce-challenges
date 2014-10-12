#include <stdio.h>

void skip(int a, FILE *fp) {
	int i;
	for (i = 0; i < a; i++)
		fgetc(fp);
}

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%c", &c) != EOF) {
		switch (c) {
		case 'z':
			printf("0");
			skip(3, fp);
			break;
		case 'o':
			printf("1");
			skip(2, fp);
			break;
		case 't':
			fscanf(fp, "%c", &c);
			if (c == 'w') {
				printf("2");
				skip(1, fp);
			} else {
				printf("3");
				skip(3, fp);
			}
			break;
		case 'f':
			fscanf(fp, "%c", &c);
			if (c == 'o') {
				printf("4");
				skip(2, fp);
			} else {
				printf("5");
				skip(2, fp);
			}
			break;
		case 's':
			fscanf(fp, "%c", &c);
			if (c == 'i') {
				printf("6");
				skip(1, fp);
			} else {
				printf("7");
				skip(3, fp);
			}
			break;
		case 'e':
			printf("8");
			skip(4, fp);
			break;
		case 'n':
			printf("9");
			skip(3, fp);
			break;
		}
		fscanf(fp, "%c", &c);
		if (c == '\n')
			printf("\n");
	}
	return 0;
}
