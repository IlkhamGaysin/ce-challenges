import fileinput

for line in fileinput.input():
    s = line.rstrip('\n').split()
    l, h = 0, int(s[0])
    for i in range(1, len(s)):
        c = (l + h) % 2
        if s[i] == "Lower":
            h = (l + h) // 2 + c - 1
        elif s[i] == "Higher":
            l = (l + h) // 2 + c + 1
        else:
            print((l + h) // 2 + c)
