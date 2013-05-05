import fileinput

for line in fileinput.input():
    st = line.split(' ')
    n = len(st) - 1
    print st[n - int(st[n])]
