#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct node {
	char		sym;
	struct node	*next;
};

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;
	struct node *head = NULL, *temp = NULL;
	bool valid = true;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		if (!valid) {
			if (c == '\n') {
				valid = true;
			}
			continue;
		}
		switch (c) {
		case '\n':
			if (head) {
				puts("False");
				while (head) {
					struct node *done = head;
					head = head->next;
					free(done);
				}
				valid = true;
			} else {
				puts("True");
			}
			break;
		case ')':
			if (head && head->sym == '(') {
				struct node *done = head;
				head = head->next;
				free(done);
			} else {
				valid = false;
			}
			break;
		case ']':
			if (head && head->sym == '[') {
				struct node *done = head;
				head = head->next;
				free(done);
			} else {
				valid = false;
			}
			break;
		case '}':
			if (head && head->sym == '{') {
				struct node *done = head;
				head = head->next;
				free(done);
			} else {
				valid = false;
			}
			break;
		default:
			temp = malloc(sizeof(struct node));
			temp->sym = c;
			temp->next = head;
			head = temp;
		}
		if (!valid) {
			puts("False");
			while (head) {
				struct node *done = head;
				head = head->next;
				free(done);
			}
		}
	}
	return 0;
}
