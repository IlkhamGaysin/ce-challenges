import fileinput

for line in fileinput.input():
    st = line.split()
    prev, count, r = "", 0, []
    for i in st:
        if i == prev:
            count += 1
        else:
            if count > 0:
                r.extend([str(count), prev])
            count, prev = 1, i
    r.extend([str(count), prev])
    print(*r, sep=" ")
