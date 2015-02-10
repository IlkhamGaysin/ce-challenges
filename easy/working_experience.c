#include <stdio.h>

void skip(int a, FILE *fp) {
	int i;
	for (i = 0; i < a; i++)
		fgetc(fp);
}

int month(char c, FILE *fp) {
	int m = 0;
	switch (c) {
	case 'J':
		fscanf(fp, "%c", &c);
		if (c == 'a') {
			m = 1;
			skip(2, fp);
		} else {
			fscanf(fp, "%c", &c);
			if (c == 'n') {
				m = 6;
			} else {
				m = 7;
			}
			skip(1, fp);
		}
		break;
	case 'F':
		m = 2;
		skip(3, fp);
		break;
	case 'M':
		skip(1, fp);
		fscanf(fp, "%c", &c);
		if (c == 'r') {
			m = 3;
		} else {
			m = 5;
		}
		skip(1, fp);
		break;
	case 'A':
		fscanf(fp, "%c", &c);
		if (c == 'p') {
			m = 4;
		} else {
			m = 8;
		}
		skip(2, fp);
		break;
	case 'S':
		m = 9;
		skip(3, fp);
		break;
	case 'O':
		m = 10;
		skip(3, fp);
		break;
	case 'N':
		m = 11;
		skip(3, fp);
		break;
	case 'D':
		m = 12;
		skip(3, fp);
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
		int i, t0, t1, m = month(c, fp) - 1;
		fscanf(fp, "%d", &t0);
		t0 = 12 * (t0 - 1990) + m;
		skip(1, fp);
		fscanf(fp, "%c", &c);
		m = month(c, fp) - 1;
		fscanf(fp, "%d", &t1);
		t1 = 12 * (t1 - 1990) + m;
		for (i = t0; i <= t1; i++) {
			work[i / 30] |= (1 << (i % 30));
		}
		fscanf(fp, "%c", &c);
		if (c == ';') {
			skip(1, fp);
		} else {
			int j, w = 0;
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
