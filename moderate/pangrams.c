#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct node {
	struct node	*next;
	char		data;
};

void pang(struct node *head);

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		struct node *head = NULL;
		while (c != '\n') {
			c = tolower(c);
			if (!islower(c))
				goto skip;

			struct node *temp, *curr = head, *oldr = NULL;
			if (head) {
				while (curr->data < c) {
					oldr = curr;
					if (!curr->next)
						break;
					curr = curr->next;
				}
				if (curr->data == c)
					goto skip;
			}
			temp = malloc(sizeof(struct node));
			temp->data = c;
			if (oldr) {
				temp->next = oldr->next;
				oldr->next = temp;
			} else {
				temp->next = head;
				head = temp;
			}
skip:
			if ((c = getc(fp)) == EOF)
				return 0;
		}
		pang(head);
		if (head) {
			struct node *curr = head, *oldr;
			while (curr) {
				oldr = curr;
				curr = curr->next;
				free(oldr);
			}
		}
	}
	return 0;
}

void pang(struct node *head)
{
	struct node *curr = head;
	bool p = true;
	char i;

	for (i = 97; i < 123; i++) {
		bool b = true;
		if (head) {
			while (curr->data < i) {
				if (!curr->next)
					break;
					curr = curr->next;
			}
			if (curr->data == i)
				b = false;
		}
		if (b) {
			printf("%c", i);
			p = false;
		}
	}
	if (p)
		printf("NULL");
	printf("\n");
}
