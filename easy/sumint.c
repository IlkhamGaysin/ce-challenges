#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
    FILE *fp;
    char line[22];
    long long a = 0;

    fp = fopen(*++argv, "r");
    while (fgets(line, 22, fp) != 0)
	a += atoll(line);
    printf("%lld\n", a);
    return 0;
}
