#include <stdbool.h>
#include <stdio.h>

bool incircle(float a, float b, float c) {
	return a*a + b*b <= c*c;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	float cx, cy, r, px, py;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "Center: (%f, %f); Radius: %f; Point: (%f, %f)\n", &cx, &cy, &r, &px, &py) != EOF) {
		if (incircle(cx-px, cy-py, r)) {
			puts("true");
		} else {
			puts("false");
		}
	}
	return 0;
}
