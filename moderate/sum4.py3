import fileinput

def numzero(s, c, z):
    if c == 0:
        return 1 if z == 0 else 0
    if len(s) < c or z + c * s[0] > 0 or z + c * s[-1] < 0:
        return 0
    return numzero(s[1:], c-1, z+s[0]) + numzero(s[1:], c, z)

for line in fileinput.input():
    print(numzero(sorted([int(i) for i in line.split(',')]), 4, 0))
