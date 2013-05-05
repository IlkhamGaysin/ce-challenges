#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct node {
	struct node	*next;
	char		data;
	bool		uniq;
};

void firstnrc(struct node *head);

int main(int argc, char *argv[])
{
	FILE *fp;
	char c;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF) {
		struct node *head = NULL;
		while (c != '\n') {
			struct node *temp;
			temp = malloc(sizeof(struct node));
			temp->data = c;
			temp->uniq = true;
			temp->next = NULL;
			if (head) {
				struct node *curr = head, *oldr;
				while (curr) {
					if (curr->data == temp->data) {
						curr->uniq = false;
						goto skip;
					}
					oldr = curr;
					curr = curr->next;
				}
				oldr->next = temp;
			} else {
				head = temp;
			}

skip:
			if ((c = getc(fp)) == EOF)
				return 0;
		}
		firstnrc(head);
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

void firstnrc(struct node *head)
{
	struct node *curr = head;

	if (head) {
		while (curr) {
			if (curr->uniq) {
				printf("%c\n", curr->data);
				return;
			}
			curr = curr->next;
		}
	}
	printf("\n");
}
