import fileinput

for line in fileinput.input():
    st = line.split(',')
    if len(st) == 2:
        print(1 if st[0].endswith(st[1].rstrip('\n')) else 0)
