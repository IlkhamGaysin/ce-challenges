import fileinput

for test in fileinput.input():
    st = test.strip()
    c = 0
    for i in st:
        c += int(i)
    print c
