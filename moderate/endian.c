#include <stdio.h>

int main() {
	int i = 1;
	char *p = (char *)&i;

	puts(*p ? "LittleEndian" : "BigEndian");

	return 0;
}
