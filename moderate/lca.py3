import fileinput

parent = {3: 8, 8: 30, 10: 20, 20: 8, 29: 20, 30: None, 52: 30}

for line in fileinput.input():
    n0, n1 = [int(i) for i in line.split()]
    found = False
    while n0 and n0 != n1:
        i = n1
        while i:
            if n0 == i:
                found = True
                break
            i = parent[i]
        if found:
            break
        n0 = parent[n0]
    print(n0)
