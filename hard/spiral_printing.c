#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int sbs = 32, s = 0, ibs = 32, i = 0, j, n, m;
	char c;
	char *sb = malloc(sbs);
	int *ib = malloc(ibs * sizeof(int));
	ib[0] = 0;

	fp = fopen(*++argv, "r");
	while ((fscanf(fp, "%d;%d;", &n, &m)) != EOF) {
		int tn = 0, te, ts, tw = 0;
		if (n * m > ibs) {
			ibs = n * m;
			ib = realloc(ib, ibs * sizeof(int));
		}
		while (i < n * m) {
			if (s == sbs) {
				sbs += sbs / 2;
				sb = realloc(sb, sbs);
			}
			if ((c = getc(fp)) == ' ' || c == '\n' || c == EOF) {
				c = '\0';
				if (++i < n * m)
					ib[i] = s + 1;
			}
			sb[s++] = c;
		}
		i = 0;
		j = 1;
		te = m - 1;
		ts = n - 1;
		printf("%s", sb);
		do {
			while (j <= te)
				printf(" %s", sb + ib[i * m + j++]);
			j--;
			i++;
			tn++;
			if (i > ts)
				break;
			while (i <= ts)
				printf(" %s", sb + ib[i++ * m + j]);
			i--;
			j--;
			te--;
			if (j < tw)
				break;
			while (j >= tw)
				printf(" %s", sb + ib[i * m + j--]);
			j++;
			i--;
			ts--;
			if (i < tn)
				break;
			while (i >= tn)
				printf(" %s", sb + ib[i-- * m + j]);
			j++;
			i++;
			tw++;
		} while (j <= te);
		printf("\n");
		s = 0;
		i = 0;
	}
	return 0;
}
