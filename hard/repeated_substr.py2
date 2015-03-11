"""find the longest repeated substring in a given text"""
import fileinput

for line in fileinput.input():
    line = line.rstrip('\n')
    l = len(line)
    if not l:
        print "NONE"
        continue
    m = ''
    n = 1
    start = 0
    while n <= ((l - start) / 2):
        f = False
        for i in range(start, l - n):
            subst = line[i : (i + n)]
            if not subst.strip():
                continue
            if subst in line[(i + n) : l]:
                m = subst
                start = i
                f = True
                break
        if not f:
            break
        n = n + 1

    print "NONE" if m == '' else m
