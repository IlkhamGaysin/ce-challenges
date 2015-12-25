#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct node {
	char		*data;
	int		length;
	struct node	*next;
};

bool leveno(char *a, char *b) {
	bool d = false;
	while (*b != '\0') {
		if (*a != *b) {
			if (d)
				return false;
			d = true;
		} else {
			b++;
		}
		a++;
	}
	return true;
}

bool levene(char *a, char *b) {
	bool d = false;
	while (*a != '\0') {
		if (*a != *b) {
			if (d)
				return false;
			d = true;
		}
		a++;
		b++;
	}
	return true;
}

int main(int argc, char *argv[]) {
	FILE *fp;
	char line[32];
	char root[6] = "hello";
	int netw = 0, strr = strlen(root);

	struct node *look = NULL, *found = NULL, *curl = NULL, *curf = NULL;
	fp = fopen(*++argv, "r");
	while (fgets(line, 31, fp) != 0) {
		int strl = strlen(line) - 1;
		struct node *temp = malloc(sizeof(struct node));
		char *w = malloc(strl);

		temp->data = strncpy(w, line, strl);
		temp->length = strl;
		temp->next = NULL;

		if ((abs(strr - strl) < 2) &&
		    ((strr == strl && levene(root, temp->data)) ||
		     (strr > strl && leveno(root, temp->data)) ||
		     (strr < strl && leveno(temp->data, root)))) {
			if (curf)
				curf->next = temp;
			else
				found = temp;
			curf = temp;
		} else {
			if (curl)
				curl->next = temp;
			else
				look = temp;
			curl = temp;
		}
	}
	struct node *done = NULL, *prev = NULL;
	curf = found;
	while (curf) {
		curl = look;
		prev = NULL;
		while (curl) {
			if ((abs(curf->length - curl->length) < 2) &&
			    ((curf->length == curl->length && levene(curf->data, curl->data)) ||
			     (curf->length > curl->length && leveno(curf->data, curl->data)) ||
			     (curf->length < curl->length && leveno(curl->data, curf->data)))) {
				if (prev) {
					prev->next = curl->next;
				} else {
					look = curl->next;
				}
				done = curl->next;
				curl->next = curf->next;
				curf->next = curl;
				curl = done;
			} else {
				prev = curl;
				curl = curl->next;
			}
		}
		done = curf;
		curf = curf->next;
		netw++;
		free(done->data);
		free(done);
	}
	printf("%d\n", netw);
	return 0;
}
