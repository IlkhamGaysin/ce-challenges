import fileinput

for line in fileinput.input():
    a = int(line)
    b = 0

    while b < 1000 and str(a) != str(a)[::-1]:
        a += int(str(a)[::-1])
        b += 1
    print str(b) + ' ' + str(a) if b < 1000 else "not found"
