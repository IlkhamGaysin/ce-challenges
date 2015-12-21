#include <stdio.h>

const char *cat[] = { "This program is for humans",
		      "Still in Mama's arms",
		      "Preschool Maniac",
		      "Elementary school",
		      "Middle school",
		      "High school",
		      "College",
		      "Working for the man",
		      "The Golden Years" };
const int age[] = { 0, 3, 5, 12, 15, 19, 23, 66, 101 };

int main(int argc, char *argv[]) {
	FILE *fp;
	int a;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		int c = 0;
		while (c < 9 && a >= age[c])
			c++;
		printf("%s\n", cat[c % 9]);
	}
	return 0;
}
