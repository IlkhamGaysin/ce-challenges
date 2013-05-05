"""Print out Pascal's triangle up to the requested depth in row major form"""
import fileinput

for line in fileinput.input():
    try:
        n = int(line)
    except ValueError:
        continue
    l = [1]
    st = [1]
    for i in xrange(n-1):
        m = [0] + l
        for jx, j in enumerate(l):
            m[jx] += j
        st += m
        l = m
    print ' '.join(str(i) for i in st)
