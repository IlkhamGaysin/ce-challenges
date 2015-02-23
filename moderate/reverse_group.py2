"""reverse the elements of the list, k items at a time"""
import fileinput

for line in fileinput.input():
    st = line.split(';')
    l = st[0].split(',')
    k = int(st[1])
    p = 0
    while p + k <= len(l):
        l = l[:p] + l[p:p+k][::-1] + l[p+k:]
        p += k
    print ','.join(l)
