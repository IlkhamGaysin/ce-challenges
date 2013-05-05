import fileinput

n = 0
for line in fileinput.input():
    if not n:
        n = int(line)
        b = [0] * n
        c = [''] * n
        continue

    l = len(line) - 1
    for ix, i in enumerate(b):
        if l > i:
            b.insert(ix, l)
            c.insert(ix, line.rstrip())
            del b[-1]
            del c[-1]
            break

for i in c:
    print i
