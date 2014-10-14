#include <stdio.h>

int main(int argc, char *argv[])
{
    FILE *fp;
    char c, line[64];
    int d;

    fp = fopen(*++argv, "r");
    while (fgets(line, 64, fp) != 0) {
        if (line[0] == '\n')
            continue;
        while (1) {
            fscanf(fp, "%c", &c);
            if (c == '\n')
                break;
            fscanf(fp, "%d", &d);
            printf("%c", line[d-1]);
        }
        printf("\n");
    }
    return 0;
}
