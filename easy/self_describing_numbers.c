#include <stdio.h>

int main(int argc, char *argv[])
{
    FILE *fp;
    char line[12];

    fp = fopen(*++argv, "r");
    while (fgets(line, 12, fp) != 0) {
	int i = 0, j = 1, a[10] = {0}, b[10] = {0};
	while (line [i] != '\n') {
	    a[i] = line[i] - 48;
	    b[a[i++]]++;
	}
	while (i-- > 0)
	    if (a[i] != b[i])
		j = 0;
	printf("%d\n", j);
    }
    return 0;
}
