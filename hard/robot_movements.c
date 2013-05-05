#include <stdio.h>
#include <stdlib.h>

struct node {
	int		data;
	struct node	*next;
};

int main()
{
	int i, res;
	struct node *head = NULL;

	for (i = 0; i < 1 << 27; i += 4) {
		int j, r = 0, c = 0, cur = i;
		int been[4][4] = { 1, 0 };

		for (j = 0; j < 15; j++) {
			switch (cur % 4) {
			case 0:
				c++;
				break;
			case 1:
				r++;
				break;
			case 2:
				c--;
				break;
			case 3:
				r--;
			}
			if (r < 0 || c < 0 || been[r][c] || r > 3 || c > 3)
				break;
			if (r == 3 && c == 3) {
				int low = i % (1 << j * 2);
				struct node *temp, *curr = head, *oldr = NULL;

				if (head) {
					while (curr->data < low) {
						oldr = curr;
						if (!curr->next)
							break;
						curr = curr->next;
					}
				if (curr->data == low)
					goto out;
				}
				temp = malloc(sizeof(struct node));
				temp->data = low;
				if (oldr) {
					temp->next = oldr->next;
					oldr->next = temp;
				} else {
					temp->next = head;
					head = temp;
				}
				res++;
out:
				break;
			}
			been[r][c] = 1;
			cur /= 4;
		}
	}
	printf("%d\n", 2 * res);
	return 0;
}
