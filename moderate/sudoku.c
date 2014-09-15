#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int sq(int a) {
    int ret = -1;
    switch (a) {
    case 4:
	ret = 2;
	break;
    case 9:
	ret = 3;
	break;
    }
    return ret;
}

int main(int argc, char *argv[])
{
    FILE *fp;
    int a;

    fp = fopen(*++argv, "r");
    while (fscanf(fp, "%d%*c", &a) != EOF) {
	bool valid = true;
	int d, i, j, crow;
	int *csqu = calloc(sq(a), sizeof(int));
	int *col = calloc(a, sizeof(int));

	for (i = 0; i < a; i++) {
	    crow = 0;
	    for (j = 0; j < a; j++) {
		fscanf(fp, "%d%*c", &d);
		d = 1 << (d - 1);
		crow |= d;
		csqu[j / sq(a)] |= d;
		col[j] |= d;
	    }
	    d = (1 << a) - 1;
	    if (crow != d) {
		valid = false;
	    }
	    if ((i + 1) % sq(a) == 0) {
		for (j = 0; j < sq(a); j++) {
		    if (csqu[j] != d) {
			valid = false;
		    } else {
			csqu[j] = 0;
		    }
		}
	    }
	}
	for (i = 0; valid && i < a; i++) {
	    if (col[i] != (1 << a) - 1) {
		valid = false;
	    }
	}
	if (valid)
	    printf("True\n");
	else
	    printf("False\n");
	free(col);
	free(csqu);
    }
    return 0;
}
