"""Determine the longest common subsequence between two strings."""
import fileinput

def btr(ccc, xxx, yyy, iii, jjj):
    """Backtrack the choices taken when computing the table."""
    if iii < 1 or jjj < 1:
        return ""
    elif xxx[iii-1] == yyy[jjj-1]:
        return btr(ccc, xxx, yyy, iii-1, jjj-1) + xxx[iii-1]
    else:
        if ccc[iii][jjj-1] > ccc[iii-1][jjj]:
            return btr(ccc, xxx, yyy, iii, jjj-1)
        else:
            return btr(ccc, xxx, yyy, iii-1, jjj)

for test in fileinput.input():
    st = test.rstrip('\n').split(';')
    try:
        x = st[0]
        y = st[1]
    except IndexError:
        continue

    c = [[0 for _ in range(len(y)+1)] for _ in range(len(x)+1)]
    for ix, i in enumerate(x):
        for jx, j in enumerate(y):
            if i == j:
                c[ix+1][jx+1] = c[ix][jx] + 1
            else:
                c[ix+1][jx+1] = max(c[ix+1][jx], c[ix][jx+1])

    print btr(c, x, y, len(x), len(y))
