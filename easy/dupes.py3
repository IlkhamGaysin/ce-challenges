import fileinput

for line in fileinput.input():
    s = line.rstrip('\n').split(',')
    print(','.join([i for ix, i in enumerate(s) if ix == 0 or i != s[ix-1]]))
