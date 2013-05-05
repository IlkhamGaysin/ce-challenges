#include <stdio.h>

int main(int argc, char *argv[])
{
    FILE *fp;
    char line[22];

    fp = fopen(*++argv, "r");
    while (fgets(line, 22, fp) != 0) {
	int i = 0, j = 0;
	while (line[i] != '\n')
	    j += line[i++] - 48;
	printf("%d\n", j);
    }
    return 0;
}
