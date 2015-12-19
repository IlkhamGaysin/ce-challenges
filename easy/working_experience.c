#include <stdio.h>

int month(char c, FILE *fp) {
	int m = 0;
	switch (c) {
	case 'J':
		fscanf(fp, "%c", &c);
		if (c == 'a') {
			m = 1;
			fseek(fp, 2, SEEK_CUR);
		} else {
			fscanf(fp, "%c", &c);
			m = c == 'n' ? 6 : 7;
			fseek(fp, 1, SEEK_CUR);
		}
		break;
	case 'F':
		m = 2;
		fseek(fp, 3, SEEK_CUR);
		break;
	case 'M':
		fseek(fp, 1, SEEK_CUR);
		fscanf(fp, "%c", &c);
		m = c == 'r' ? 3 : 5;
		fseek(fp, 1, SEEK_CUR);
		break;
	case 'A':
		fscanf(fp, "%c", &c);
		m = c == 'p' ? 4 : 8;
		fseek(fp, 2, SEEK_CUR);
		break;
	case 'S':
		m = 9;
		fseek(fp, 3, SEEK_CUR);
		break;
	case 'O':
		m = 10;
		fseek(fp, 3, SEEK_CUR);
		break;
	case 'N':
		m = 11;
		fseek(fp, 3, SEEK_CUR);
		break;
	case 'D':
		m = 12;
		fseek(fp, 3, SEEK_CUR);
		break;
	}
	return m;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;
	int work[12] = { 0 };

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%c", &c) != EOF) {
		unsigned i, t0, t1, m = month(c, fp) - 1;
		fscanf(fp, "%d", &t0);
		t0 = 12 * (t0 - 1990) + m;
		fseek(fp, 1, SEEK_CUR);
		fscanf(fp, "%c", &c);
		m = month(c, fp) - 1;
		fscanf(fp, "%d", &t1);
		t1 = 12 * (t1 - 1990) + m;
		for (i = t0; i <= t1; i++)
			work[i / 30] |= (1 << (i % 30));
		fscanf(fp, "%c", &c);
		if (c == ';') {
			fseek(fp, 1, SEEK_CUR);
		} else {
			unsigned j, w = 0;
			for (i = 0; i < 12; i++) {
				for (j = 0; j < 30; j++)
					if (work[i] & (1 << j))
						w++;
				work[i] = 0;
			}
			printf("%d\n", w / 12);
		}
	}
	return 0;
}
