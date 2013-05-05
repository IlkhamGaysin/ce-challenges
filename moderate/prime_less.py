from __future__ import print_function
import fileinput

for line in fileinput.input():
    n = int(line)
    if n > 2:
        print(2, end = '')
    s = [True] * (n / 2)
    for i in xrange(3, int(n ** 0.5) + 1, 2):
        if s[i / 2]:
            s[i * i / 2::i] = [False] * ((n - i * i - 1) / (2 * i) + 1)
    for i in xrange(1, n / 2):
        if s[i]:
            print(',', 2 * i + 1, sep='', end = '')
    print()
