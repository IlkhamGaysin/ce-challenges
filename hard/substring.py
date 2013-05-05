"""Determine if the second string is a substring of the first"""
import fileinput

def subcheck(qqq, rrr):
    """must not use any substr type library function"""
    while qqq < len(st[0]) and rrr < len(st[1]):
        if st[1][rrr] == '*':
            mmm = False
            for iii in xrange(qqq, len(st[0])):
                if subcheck(iii, rrr+1):
                    mmm = True
                    break
            return mmm
        elif st[1][rrr] == '\\' and rrr < len(st[1])+1 and st[1][rrr+1] == '*':
            if st[0][qqq] == '*':
                qqq += 1
                rrr += 2
                continue
            else:
                return False
        if st[0][qqq] != st[1][rrr]:
            return False
        qqq += 1
        rrr += 1
    if rrr < len(st[1]):
        return False
    return True

for line in fileinput.input():
    st = line.split(',')
    st[1] = st[1].lstrip('*')
    st[1] = st[1].rstrip('\n')
    while len(st[1]) > 1 and st[1][-1] == '*' and st[1][-2] != '\\':
        st[1] = st[1][:-2]

    match = False
    for i in xrange(len(st[0])):
        if subcheck(i, 0):
            match = True
            break
    if match:
        print "true"
    else:
        print "false"
