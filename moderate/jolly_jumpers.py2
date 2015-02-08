import fileinput

for line in fileinput.input():
    st = line.split(' ')
    if st.pop(0) == '1':
        print "Jolly"
        continue
    s = []
    for i in xrange(len(st) - 1):
        s.append(abs(int(st[i]) - int(st[i+1])))
    print "Jolly" if sorted(s) == range(1, len(st)) else "Not jolly"
