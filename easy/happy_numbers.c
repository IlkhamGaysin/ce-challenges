#include <stdio.h>
#include <stdlib.h>

int happy(int a)
{
	int i = 0, ret = 0;
	char line[13];

	sprintf(line, "%d", a);
	do {
		int c = line[i] - 48;
		ret += c * c;
	} while (line[++i] != '\0');
	return ret;
}

int main(int argc, char *argv[])
{
    FILE *fp;
    char line[22];

    fp = fopen(*++argv, "r");
    while (fgets(line, 22, fp) != 0) {
	int i = 0, a = atoi(line), b[128] = {a,0};

	while (a != 1) {
	    int j = 0;
	    if (i > 126)
		goto out;
	    a = happy(a);
	    while (b[j] != 0) {
		if (b[j++] == a)
		    goto out;
	    }
	    b[++i] = a;
	}
	printf("1\n");
	continue;
out:
	printf("0\n");
    }
    return 0;
}
