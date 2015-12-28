#include <stdbool.h>
#include <stdio.h>

static bool incircle(float a, float b, float c) {
	return a * a + b * b <= c * c;
}

int main(int argc, char *argv[]) {
	FILE *fp;
	float cx, cy, r, px, py;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "Center: (%f, %f); Radius: %f; Point: (%f, %f)\n", &cx, &cy, &r, &px, &py) != EOF)
		puts(incircle(cx - px, cy - py, r) ? "true" : "false");
	return 0;
}
