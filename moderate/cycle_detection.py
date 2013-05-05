import fileinput

for line in fileinput.input():
    st = line.split()

    f = []
    j = {}
    while st:
        i = st.pop(0)
        if i not in j:
            if not st:
                break
            j[i] = st[0]
            if f:
                break
        else:
            if i in f:
                break
            f.append(i)
    if f:
        print ' '.join(i for i in f)
